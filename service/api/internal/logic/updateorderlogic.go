package logic

import (
	"context"
	"database/sql"
	"time"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/customer"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UpdateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderLogic) UpdateOrder(req *types.UpdateOrderRequest) (resp *types.UpdateOrderResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role == variables.Company || role == variables.Employee {
		return nil, status.Errorf(401, "Invalid, Not customer.")
	}

	// Get origin order
	res, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(404, "Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	reserve_date, err := time.Parse("2006-01-02 15:04:05", req.Reserve_date)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// if modify reserve_date
	if reserve_date != res.ReserveDate {
		// if reserve_date is close to current time in 12 hours, return error
		if time.Now().Add(time.Hour * 12).Before(reserve_date) {
			return nil, status.Error(500, "Reserve date is futher less than 12 hours.")
		}
		res.ReserveDate = reserve_date
	}
	res.OrderDescription = sql.NullString{req.Order_description, req.Order_description != ""}

	// Modify address details
	err = l.svcCtx.BAddressModel.Update(l.ctx, &address.BAddress{
		AddressDetails: req.Address_info.Address_details,
		Suburb:         req.Address_info.Suburb,
		Postcode:       req.Address_info.Postcode,
		StateCode:      req.Address_info.State_code,
		Country:        sql.NullString{req.Address_info.Country, req.Address_info.Country != ""},
	})
	if err != nil {
		if err == address.ErrNotFound {
			return nil, status.Error(404, "Invalid, Address not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Modify customer details
	err = l.svcCtx.BCustomerModel.Update(l.ctx, &customer.BCustomer{
		CustomerId:     uid,
		CustomerName:   req.Customer_info.Customer_name,
		CountryCode:    req.Customer_info.Country_code,
		ContactDetails: req.Customer_info.Contact_details,
	})

	if err != nil {
		if err == customer.ErrNotFound {
			return nil, status.Error(404, "Invalid, Customer not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.BOrderModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &types.UpdateOrderResponse{}, nil
}
