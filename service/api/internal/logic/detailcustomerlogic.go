package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/customer"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DetailCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailCustomerLogic {
	return &DetailCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailCustomerLogic) DetailCustomer(req *types.DetailCustomerRequest) (resp *types.DetailCustomerResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	var res *customer.BCustomer
	if role == variables.Customer {
		res, err = l.svcCtx.BCustomerModel.FindOne(l.ctx, uid)
	} else {
		res, err = l.svcCtx.BCustomerModel.FindOne(l.ctx, req.Customer_id)
	}

	if err != nil {
		if err == customer.ErrNotFound {
			return nil, status.Error(404, "Invalid, Customer not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &types.DetailCustomerResponse{
		Customer_id:     res.CustomerId,
		Customer_name:   res.CustomerName,
		Country_code:    res.CountryCode,
		Contact_details: res.ContactDetails,
	}, nil
}
