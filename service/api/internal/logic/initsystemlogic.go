package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitSystemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitSystemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitSystemLogic {
	return &InitSystemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitSystemLogic) InitSystem(req *types.InitSystemRequest) (resp *types.InitSystemResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
