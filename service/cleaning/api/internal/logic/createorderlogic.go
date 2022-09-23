package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	"cleaningservice/common/errorx"
	"cleaningservice/common/orderqueue"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/address"
	"cleaningservice/service/cleaning/model/broadcast"
	"cleaningservice/service/cleaning/model/customer"
	"cleaningservice/service/cleaning/model/order"
	"cleaningservice/service/cleaning/model/orderqueue/awaitqueue"
	"cleaningservice/service/cleaning/model/payment"
	"cleaningservice/service/cleaning/model/property"
	"cleaningservice/service/cleaning/model/region"
	"cleaningservice/service/cleaning/model/service"
	"cleaningservice/service/cleaning/validation"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderRequest) (resp *types.CreateOrderResponse, err error) {
	logx.Info("function entrance here\n")
	// Exist detail check and create if new details

	// Payment
	// expiry time convert
	exp_time, err := time.Parse("2006-01-02 15:04:05", req.Deposite_info.Expiry_time)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}
	// payment structure
	payment_item := &payment.BPayment{
		CardNumber:   req.Deposite_info.Card_number,
		HolderName:   req.Deposite_info.Holder_name,
		ExpiryTime:   exp_time,
		SecurityCode: req.Deposite_info.Security_code,
	}

	_, err = l.svcCtx.BPaymentModel.Enquire(l.ctx, payment_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Customer
	// Check if contact details valid
	if !validation.CheckCustomerPhone(req.Customer_info.Customer_phone) {
		return nil, errorx.NewCodeError(500, err.Error())
	}
	if !validation.CheckCustomerEmail(req.Customer_info.Customer_email) {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// customer structure
	customer_item := &customer.BCustomer{
		CustomerName:  req.Customer_info.Customer_name,
		CountryCode:   "61",
		CustomerPhone: req.Customer_info.Customer_phone,
		CustomerEmail: req.Customer_info.Customer_email,
	}

	customer_item, err = l.svcCtx.BCustomerModel.Enquire(l.ctx, customer_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Region
	region_item := &region.BRegion{
		RegionName: req.Address_info.Suburb,
		RegionType: "Suburb",
		Postcode:   req.Address_info.Postcode,
		StateCode:  req.Address_info.State_code,
		StateName:  req.Address_info.State_name,
	}

	// Address
	address_item := &address.BAddress{
		Street:    req.Address_info.Street,
		Suburb:    req.Address_info.Suburb,
		Postcode:  req.Address_info.Postcode,
		Property:  strings.ToLower(req.Address_info.Property),
		City:      req.Address_info.City,
		Lat:       req.Address_info.Lat,
		Lng:       req.Address_info.Lng,
		Formatted: req.Address_info.Formatted,
	}

	// Check if data valid
	if ret, err := validation.CheckAddressDetails(address_item); !ret {
		return nil, errorx.NewCodeError(500, "Invalid address details"+err.Error())
	}

	// Enquire region
	region_item, err = l.svcCtx.BRgionModel.Enquire(l.ctx, region_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}
	// Enquire address
	err = l.svcCtx.BAddressModel.Enquire(l.ctx, address_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Get Property details
	property_item, err := l.svcCtx.BPorpertyModel.FindOneByPropertyName(l.ctx, address_item.Property)
	if err != nil {
		if err == property.ErrNotFound {
			return nil, status.Error(404, "Invalid, Property type not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Get category info for emailing
	_, err = l.svcCtx.BCategoryModel.FindOne(l.ctx, req.Category_id)
	if err != nil {
		if err == service.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Category not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Property and region charge
	pr_charge := float64(1.0 + float64(region_item.ChargeAmount+property_item.ChargeAmount)/100.0)

	// Service
	// Basic Services
	var item_amount float64 = 0
	for index, basic_service := range req.Basic_items.Items {
		service_item, err := l.svcCtx.BServiceModel.FindOne(l.ctx, basic_service.Service_id)
		if err != nil {
			if err == service.ErrNotFound {
				return nil, errorx.NewCodeError(404, "Invalid, Basic service(s) not found.")
			}
			return nil, errorx.NewCodeError(500, err.Error())
		}
		if req.Category_id != service_item.ServiceType {
			logx.Alert("Wrong service is intended to be implemented.")
			continue
		}
		// Calculate fees and get full service items
		item_amount += service_item.ServicePrice * float64(basic_service.Service_quantity)

		// Get basic items details
		req.Basic_items.Items[index].Service_name = service_item.ServiceName
		req.Basic_items.Items[index].Service_scope = service_item.ServiceScope
		req.Basic_items.Items[index].Service_price = service_item.ServicePrice * pr_charge
	}

	basic_items, err := json.Marshal(req.Basic_items)
	if err != nil {
		return nil, errorx.NewCodeError(404, "Invalid, Basic service(s) marshal failed.")
	}

	// Additional Services
	for index, addition_service := range req.Additional_items.Items {
		service_item, err := l.svcCtx.BServiceModel.FindOne(l.ctx, addition_service.Service_id)
		if err != nil {
			if err == service.ErrNotFound {
				return nil, errorx.NewCodeError(404, "Invalid, Additional service(s) not found.")
			}
			return nil, errorx.NewCodeError(500, err.Error())
		}

		if req.Category_id != service_item.ServiceType {
			logx.Alert("Wrong service is intended to be implemented.")
			continue
		}
		item_amount += service_item.ServicePrice * float64(addition_service.Service_quantity)

		// Get additional items details
		req.Additional_items.Items[index].Service_name = service_item.ServiceName
		req.Additional_items.Items[index].Service_price = service_item.ServicePrice * pr_charge
		req.Additional_items.Items[index].Service_scope = service_item.ServiceScope
	}

	additional_items, err := json.Marshal(req.Additional_items)
	if err != nil {
		return nil, errorx.NewCodeError(404, "Invalid, Additional service(s) marshal failed.")
	}

	// Calculate fees
	gst_amount := item_amount / variables.GST
	total_amount := item_amount + gst_amount
	deposite_amount := total_amount / variables.Deposite_rate
	final_amount := total_amount - deposite_amount

	// Order
	// Convert reserve date time
	reserve_date, err := time.Parse("2006-01-02 15:04:05", req.Reserve_date)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Create order
	order_item := order.BOrder{
		CustomerId:          customer_item.CustomerId,
		AddressId:           address_item.AddressId,
		ContractorId:        sql.NullInt64{Int64: 0, Valid: false},
		FinanceId:           sql.NullInt64{Int64: 0, Valid: false},
		CategoryId:          req.Category_id,
		BasicItems:          string(basic_items),
		AdditionalItems:     sql.NullString{String: string(additional_items), Valid: len(req.Additional_items.Items) > 0},
		OrderDescription:    sql.NullString{String: req.Order_description, Valid: true},
		CurrentDepositeRate: int64(variables.Deposite_rate),
		DepositeAmount:      deposite_amount,
		FinalAmount:         final_amount,
		ItemAmount:          item_amount,
		GstAmount:           gst_amount,
		TotalAmount:         total_amount,
		BalanceAmount:       0,
		PostDate:            time.Now(),
		ReserveDate:         reserve_date,
		FinishDate:          sql.NullTime{Time: time.Now(), Valid: false},
		PaymentDate:         sql.NullTime{Time: time.Now(), Valid: false},
		Status:              order.Queuing,
		UrgantFlag:          0,
	}

	order_res, err := l.svcCtx.BOrderModel.Insert(l.ctx, &order_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}
	newId, err := order_res.LastInsertId()
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Timing to broadcast the order
	l.broadcastOrder(newId, req.Category_id)

	return &types.CreateOrderResponse{
		Order_id: newId,
	}, nil
}

func (l *CreateOrderLogic) broadcastOrder(orderId int64, categoryId int64) {
	go l.svcCtx.BBroadcastModel.Insert(&broadcast.BBroadcast{
		GroupId: categoryId,
		OrderId: orderId,
	})

	go orderqueue.PushOne(orderId)
	go l.svcCtx.RAwaitQueueModel.Insert(&awaitqueue.RAwaitQueue{OrderId: orderId, Vacancy: 0})
}
