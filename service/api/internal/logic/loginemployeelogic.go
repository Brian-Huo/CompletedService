package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginEmployeeLogic {
	return &LoginEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginEmployeeLogic) LoginEmployee(req *types.LoginEmployeeRequest) (resp *types.LoginEmployeeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
