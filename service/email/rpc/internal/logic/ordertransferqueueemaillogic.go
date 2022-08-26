package logic

import (
	"context"

	"cleaningservice/service/email/rpc/internal/svc"
	"cleaningservice/service/email/rpc/types/email"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderTransferQueueEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderTransferQueueEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderTransferQueueEmailLogic {
	return &OrderTransferQueueEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderTransferQueueEmailLogic) OrderTransferQueueEmail(in *email.OrderTransferQueueEmailRequest) (*email.OrderTransferQueueEmailResponse, error) {
	// todo: add your logic here and delete this line

	return &email.OrderTransferQueueEmailResponse{}, nil
}
