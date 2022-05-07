package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginServiceLogic {
	return &LoginServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginServiceLogic) LoginService(req *types.LoginServiceRequest) (resp *types.LoginServiceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
