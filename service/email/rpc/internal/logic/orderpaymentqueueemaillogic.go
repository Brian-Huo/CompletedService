package logic

import (
	"context"

	"cleaningservice/service/email/rpc/internal/svc"
	"cleaningservice/service/email/rpc/types/email"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPaymentQueueEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderPaymentQueueEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPaymentQueueEmailLogic {
	return &OrderPaymentQueueEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderPaymentQueueEmailLogic) OrderPaymentQueueEmail(in *email.OrderPaymentQueueEmailRequest) (*email.OrderPaymentQueueEmailResponse, error) {
	// todo: add your logic here and delete this line

	return &email.OrderPaymentQueueEmailResponse{}, nil
}
