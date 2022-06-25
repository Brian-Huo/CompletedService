package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfirmOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfirmOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmOrderLogic {
	return &ConfirmOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfirmOrderLogic) ConfirmOrder(req *types.ConfirmOrderRequest) (resp *types.ConfirmOrderResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
