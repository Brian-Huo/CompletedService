package logic

import (
	"context"
	"database/sql"
	"time"

	"cleaningservice/common/broadcast"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/customer"
	"cleaningservice/service/model/order"
	"cleaningservice/service/model/payment"
	"cleaningservice/service/model/service"

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
	// Exist detail check and create if new details
	// Payment
	var paymentId int64
	payment_item, err := l.svcCtx.BPaymentModel.FindOneByCardNumber(l.ctx, req.Deposite_info.Card_number)
	if err == payment.ErrNotFound {
		// expiry time convert
		exp_time, err := time.Parse("2006-01-02 15:04:05", req.Deposite_info.Expiry_time)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		newPayment := payment.BPayment{
			CardNumber:   req.Deposite_info.Card_number,
			HolderName:   req.Deposite_info.Holder_name,
			ExpiryTime:   exp_time,
			SecurityCode: req.Deposite_info.Security_code,
		}

		res, err := l.svcCtx.BPaymentModel.Insert(l.ctx, &newPayment)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		paymentId, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	} else if err == nil {
		paymentId = payment_item.PaymentId
	} else {
		return nil, status.Error(500, err.Error())
	}

	// Customer
	var customerId int64
	customer_item, err := l.svcCtx.BCustomerModel.FindOneByContactDetails(l.ctx, req.Customer_info.Contact_details)
	if err == customer.ErrNotFound {
		newCustomer := customer.BCustomer{
			CustomerName:   req.Customer_info.Customer_name,
			CountryCode:    req.Customer_info.Country_code,
			ContactDetails: req.Customer_info.Contact_details,
		}

		res, err := l.svcCtx.BCustomerModel.Insert(l.ctx, &newCustomer)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		customerId, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	} else if err == nil {
		customerId = customer_item.CustomerId
	} else {
		return nil, status.Error(500, err.Error())
	}

	// Address
	address_Item := address.BAddress{
		Street:    req.Address_info.Street,
		Suburb:    req.Address_info.Suburb,
		Postcode:  req.Address_info.Postcode,
		StateCode: req.Address_info.State_code,
		Country:   "AU",
	}

	res, err := l.svcCtx.BAddressModel.Insert(l.ctx, &address_Item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	addressId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Get time variables
	reserve_date, err := time.Parse("2006-01-02 15:04:05", req.Reserve_date)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Calculate fees and get full service strings
	var service_fee float64 = 0
	var service_list string = ""
	for _, service_id := range req.Service_list {
		service_list += variables.Separator
		service_item, err := l.svcCtx.BServiceModel.FindOne(l.ctx, service_id)
		if err != nil {
			if err == service.ErrNotFound {
				return nil, status.Error(404, "Invalid, Service(s) not found.")
			}
			return nil, status.Error(500, err.Error())
		}
		service_fee += service_item.ServicePrice
		service_list += service_item.ServiceType
	}

	deposite_amount := service_fee / variables.Deposite_rate
	final_amount := service_fee - deposite_amount
	gst_amount := service_fee / variables.GST
	total_fee := service_fee + gst_amount

	// Create order
	newItem := order.BOrder{
		CustomerId:          customerId,
		AddressId:           addressId,
		CompanyId:           sql.NullInt64{0, false},
		EmployeeId:          sql.NullInt64{0, false},
		ServiceList:         service_list,
		DepositePayment:     paymentId,
		DepositeAmount:      deposite_amount,
		CurrentDepositeRate: int64(variables.Deposite_rate),
		DepositeDate:        time.Now(),
		FinalPayment:        sql.NullInt64{0, false},
		FinalAmount:         final_amount,
		FinalPaymentDate:    sql.NullTime{time.Now(), false},
		GstAmount:           gst_amount,
		TotalFee:            total_fee,
		OrderDescription:    sql.NullString{req.Order_description, true},
		PostDate:            time.Now(),
		ReserveDate:         reserve_date,
		FinishDate:          sql.NullTime{time.Now(), false},
		Status:              int64(variables.Queuing),
	}

	res, err = l.svcCtx.BOrderModel.Insert(l.ctx, &newItem)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Timing to broadcast the order
	go broadcast.TimerBroadcast()

	return &types.CreateOrderResponse{
		Order_id: newId,
	}, nil
}
