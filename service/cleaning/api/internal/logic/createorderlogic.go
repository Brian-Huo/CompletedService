package logic

import (
	"context"
	"database/sql"
	"strconv"
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
	"cleaningservice/service/cleaning/model/payment"
	"cleaningservice/service/cleaning/model/service"
	"cleaningservice/service/cleaning/validation"
	"cleaningservice/service/email/rpc/types/email"

	"github.com/zeromicro/go-zero/core/logx"
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

	payment_item, err = l.svcCtx.BPaymentModel.Enquire(l.ctx, payment_item)
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

	// Get customer info for emailing
	customer_email := email.CustomerMsg{
		CustomerId:    customer_item.CustomerId,
		CustomerName:  customer_item.CustomerName,
		CustomerType:  customer_item.CustomerType,
		CountryCode:   customer_item.CountryCode,
		CustomerPhone: customer_item.CustomerPhone,
		CustomerEmail: customer_item.CustomerEmail,
	}

	// Address
	address_item := address.BAddress{
		Street:    req.Address_info.Street,
		Suburb:    req.Address_info.Suburb,
		Postcode:  req.Address_info.Postcode,
		City:      req.Address_info.City,
		StateCode: req.Address_info.State_code,
		Country:   req.Address_info.Country,
		Lat:       req.Address_info.Lat,
		Lng:       req.Address_info.Lng,
		Formatted: req.Address_info.Formatted,
	}

	// Check if data valid
	if ret, err := validation.CheckAddressDetails(&address_item); !ret {
		return nil, errorx.NewCodeError(500, "Invalid address details"+err.Error())
	}
	address_res, err := l.svcCtx.BAddressModel.Insert(l.ctx, &address_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}
	address_item.AddressId, err = address_res.LastInsertId()
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Get address info for emailing
	address_email := email.AddressMsg{
		AddressId: address_item.AddressId,
		Street:    address_item.Street,
		Suburb:    address_item.Suburb,
		Postcode:  address_item.Postcode,
		City:      address_item.City,
		StateCode: address_item.StateCode,
		Country:   address_item.Country,
		Formatted: address_item.Formatted,
	}

	// Get category info for emailing
	category_item, err := l.svcCtx.BCategoryModel.FindOne(l.ctx, req.Category_id)
	if err != nil {
		if err == service.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Category not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Get category info for emailing
	category_email := email.CategoryMsg{
		CategoryId:          category_item.CategoryId,
		CategoryAbbr:        category_item.CategoryAddr,
		CategoryName:        category_item.CategoryName,
		CategoryDescription: category_item.CategoryDescription,
	}

	// Service
	// Get services info for emailing
	var service_email []*email.ServiceMsg

	// Calculate fees and get full service strings
	var service_fee float64 = 0
	var service_list string = ""
	for _, order_service := range req.Service_list {
		service_list += variables.Separator
		service_item, err := l.svcCtx.BServiceModel.FindOne(l.ctx, order_service.Service_id)
		if err != nil {
			if err == service.ErrNotFound {
				return nil, errorx.NewCodeError(404, "Invalid, Service(s) not found.")
			}
			return nil, errorx.NewCodeError(500, err.Error())
		}
		service_fee += service_item.ServicePrice * float64(order_service.Service_quantity)
		service_list += service_item.ServiceName + ":x" + strconv.Itoa(order_service.Service_quantity)

		// Get service info for emailing
		service_email = append(service_email, &email.ServiceMsg{
			ServiceId:          service_item.ServiceId,
			ServiceScope:       service_item.ServiceScope,
			ServiceName:        service_item.ServiceName,
			ServiceDescription: service_item.ServiceDescription,
			ServiceQuantity:    int32(order_service.Service_quantity),
			ServicePrice:       service_item.ServicePrice,
		})
	}
	service_list = strings.Replace(service_list, variables.Separator, "", 1)

	// Calculate fees
	item_amount := service_fee
	gst_amount := service_fee / variables.GST
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
		ContractorId:        sql.NullInt64{0, false},
		FinanceId:           sql.NullInt64{0, false},
		CategoryId:          req.Category_id,
		ServiceList:         service_list,
		DepositePayment:     sql.NullInt64{payment_item.PaymentId, payment_item.PaymentId != 0},
		DepositeAmount:      deposite_amount,
		DepositeDate:        sql.NullTime{time.Now(), payment_item.PaymentId != 0},
		FinalPayment:        sql.NullInt64{0, false},
		FinalAmount:         final_amount,
		FinalPaymentDate:    sql.NullTime{time.Now().Add(time.Hour * 168), false},
		CurrentDepositeRate: int64(variables.Deposite_rate),
		ItemAmount:          item_amount,
		GstAmount:           gst_amount,
		TotalAmount:         total_amount,
		OrderDescription:    sql.NullString{req.Order_description, true},
		PostDate:            time.Now(),
		ReserveDate:         reserve_date,
		FinishDate:          sql.NullTime{time.Now(), false},
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

	// Get order info for emailing
	order_email := email.OrderMsg{
		OrderId:        newId,
		DepositeAmount: order_item.DepositeAmount,
		FinalAmount:    order_item.FinalAmount,
		DepositeRate:   int32(order_item.CurrentDepositeRate),
		GstAmount:      order_item.GstAmount,
		TotalAmount:    order_item.TotalAmount,
		ReserveDate:    order_item.ReserveDate.Format("02/01/2006 15:04:05"),
	}

	// Send order Invoice
	l.svcCtx.EmailRpc.InvoiceEmail(l.ctx, &email.InvoiceEmailRequest{
		AddressInfo:  &address_email,
		CategoryInfo: &category_email,
		CustomerInfo: &customer_email,
		ServiceInfo:  service_email,
		OrderInfo:    &order_email,
	})

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

	go orderqueue.Insert(orderId)
}
