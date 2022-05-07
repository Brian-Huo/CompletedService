package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return
}
