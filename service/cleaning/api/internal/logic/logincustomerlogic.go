package logic

import (
	"context"
	"time"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/customer"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type LoginCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginCustomerLogic {
	return &LoginCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginCustomerLogic) LoginCustomer(req *types.LoginCustomerRequest) (resp *types.LoginCustomerResponse, err error) {
	// find customer by contact_details
	res, err := l.svcCtx.BCustomerModel.FindOneByContactDetails(l.ctx, req.Contact_details)
	if err == customer.ErrNotFound {
		// if customer not found, insert customer
		customer_item, err := l.svcCtx.BCustomerModel.Insert(l.ctx, &customer.BCustomer{
			CustomerName:   req.Contact_details,
			CountryCode:    "Astralia",
			ContactDetails: req.Contact_details,
		})
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		// get customer id
		newId, err := customer_item.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		res.CustomerId = newId
	} else if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 签发 jwt token
	now := time.Now().Unix()
	token, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire,
		res.CustomerId, variables.Customer)
	if err != nil {
		return nil, status.Error(500, "Jwt token error.")
	}

	return &types.LoginCustomerResponse{
		Code:        200,
		Msg:         "success",
		AccessToken: token,
	}, nil
}
