package logic

import (
	"context"
	"database/sql"
	"time"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/address"
	"cleaningservice/service/cleaning/model/customer"

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
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	} else if role == variables.Company || role == variables.Contractor {
		return nil, status.Errorf(401, "Invalid, Not customer.")
	}

	// Get origin order
	order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorx.NewCodeError(404, "Order not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	reserve_date, err := time.Parse("2006-01-02 15:04:05", req.Reserve_date)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// if modify reserve_date
	if reserve_date != order_item.ReserveDate {
		// if reserve_date is close to current time in 12 hours, return error
		if time.Now().Add(time.Hour * 12).Before(reserve_date) {
			return nil, errorx.NewCodeError(500, "Reserve date is futher less than 12 hours.")
		}
		order_item.ReserveDate = reserve_date
	}

	// if modify address
	address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, req.Address_info.Address_id)
	if err != nil {
		if err == address.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Address not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	if req.Address_info.Street != address_item.Street {
		// if address is close to current time in 12 hours, return error
		if time.Now().Add(time.Hour * 12).Before(reserve_date) {
			return nil, errorx.NewCodeError(500, "Address update is futher less than 12 hours.")
		}

		// Modify address details
		err = l.svcCtx.BAddressModel.Update(l.ctx, &address.BAddress{
			Street:    req.Address_info.Street,
			Suburb:    req.Address_info.Suburb,
			Postcode:  req.Address_info.Postcode,
			City:      req.Address_info.City,
			StateCode: req.Address_info.State_code,
			Country:   req.Address_info.Country,
			Lat:       req.Address_info.Lat,
			Lng:       req.Address_info.Lng,
			Formatted: req.Address_info.Formatted,
		})
		if err != nil {
			return nil, errorx.NewCodeError(500, err.Error())
		}
	}
	order_item.OrderDescription = sql.NullString{req.Order_description, req.Order_description != ""}

	// Modify customer details
	err = l.svcCtx.BCustomerModel.Update(l.ctx, &customer.BCustomer{
		CustomerId:     uid,
		CustomerName:   req.Customer_info.Customer_name,
		CountryCode:    req.Customer_info.Country_code,
		ContactDetails: req.Customer_info.Contact_details,
	})
	if err != nil {
		if err == customer.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Customer not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Update order
	err = l.svcCtx.BOrderModel.Update(l.ctx, order_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}
	return &types.UpdateOrderResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
