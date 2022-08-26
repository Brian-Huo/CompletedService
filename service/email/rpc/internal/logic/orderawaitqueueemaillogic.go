package logic

import (
	"context"

	"cleaningservice/service/email/rpc/internal/svc"
	"cleaningservice/service/email/rpc/types/email"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderAwaitQueueEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderAwaitQueueEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderAwaitQueueEmailLogic {
	return &OrderAwaitQueueEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderAwaitQueueEmailLogic) OrderAwaitQueueEmail(in *email.OrderAwaitQueueEmailRequest) (*email.OrderAwaitQueueEmailResponse, error) {
	// todo: add your logic here and delete this line

	return &email.OrderAwaitQueueEmailResponse{}, nil
}
