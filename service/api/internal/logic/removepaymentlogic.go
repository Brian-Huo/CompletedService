package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemovePaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemovePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemovePaymentLogic {
	return &RemovePaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemovePaymentLogic) RemovePayment(req *types.RemovePaymentRequest) (resp *types.RemovePaymentResponse, err error) {
	return nil, status.Error(500, err.Error())
}
