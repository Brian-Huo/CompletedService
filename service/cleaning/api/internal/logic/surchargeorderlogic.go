package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SurchargeOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSurchargeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SurchargeOrderLogic {
	return &SurchargeOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SurchargeOrderLogic) SurchargeOrder(req *types.SurchargeOrderRequest) (resp *types.SurchargeOrderResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
