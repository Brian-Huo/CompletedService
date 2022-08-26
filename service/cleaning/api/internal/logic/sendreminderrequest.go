package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendReminderRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendReminderRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendReminderRequestLogic {
	return &SendReminderRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendReminderRequestLogic) SendReminderRequest() (err error) {
	logx.Info("function entrance here\n")

	// // Get order queues
	// awaitqueue, err := l.svcCtx.RAwaitQueueModel.List()
	// paymentqueue, err := l.svcCtx.RPaymentQueueModel.List()
	// transferqueue, err := l.svcCtx.RTransferQueueModel.List()

	return nil
}
