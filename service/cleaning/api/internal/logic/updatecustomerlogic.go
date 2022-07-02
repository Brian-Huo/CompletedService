package logic

import (
	"context"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/customer"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomerLogic {
	return &UpdateCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCustomerLogic) UpdateCustomer(req *types.UpdateCustomerRequest) (resp *types.UpdateCustomerResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	} else if role != variables.Customer {
		return nil, errorx.NewCodeError(401, "Invalid, Not customer.")
	}

	err = l.svcCtx.BCustomerModel.Update(l.ctx, &customer.BCustomer{
		CustomerId:    uid,
		CustomerName:  req.Customer_name,
		CountryCode:   req.Country_code,
		CustomerPhone: req.Customer_phone,
		CustomerEmail: req.Customer_email,
	})

	if err != nil {
		if err == customer.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Customer not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	return &types.UpdateCustomerResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
