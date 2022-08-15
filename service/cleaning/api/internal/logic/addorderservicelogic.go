package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrderServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderServiceLogic {
	return &AddOrderServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrderServiceLogic) AddOrderService(req *types.AddOrderServiceRequest) (resp *types.AddOrderServiceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
