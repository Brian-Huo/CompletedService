package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/model/orderqueue/awaitqueue"
	"cleaningservice/service/cleaning/model/orderqueue/paymentqueue"
	"cleaningservice/service/email/rpc/types/email"

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

	// Get order queues
	awaitOrderQueue, err := l.svcCtx.RAwaitQueueModel.List()
	if err != nil {
		if err != awaitqueue.ErrNotFound {
			logx.Alert("Await order queue failed.")
			return err
		}
	} else {
		// Check await order vacancy
		err = l.svcCtx.RAwaitQueueModel.Count()
		if err != nil {
			logx.Alert("Await order queue failed.")
			return err
		}
	}

	paymentOrderQueue, err := l.svcCtx.RPaymentQueueModel.List()
	if err != nil {
		if err != paymentqueue.ErrNotFound {
			logx.Alert("Payment order queue failed.")
			return err
		}
	}

	// Send awaiting order reminder email
	var awaitOrderIds []string
	var awaitOrderVacancy []string
	for k, v := range *awaitOrderQueue {
		awaitOrderIds = append(awaitOrderIds, k)
		awaitOrderVacancy = append(awaitOrderVacancy, v)
	}

	awaitRes, err := l.svcCtx.EmailRpc.OrderAwaitQueueEmail(l.ctx, &email.OrderAwaitQueueEmailRequest{OrderId: awaitOrderIds, Vacancy: awaitOrderVacancy})
	if err != nil {
		logx.Alert("Send Daily awaiting order reminder email failed" + awaitRes.Msg)
		return err
	} else {
		logx.Info("Daily awaiting order reminder email sent.")
	}

	// Send awaiting order reminder email
	var paymentOrderIds []string
	var paymentOrderContacts []string
	for k, v := range *paymentOrderQueue {
		paymentOrderIds = append(paymentOrderIds, k)
		paymentOrderContacts = append(paymentOrderContacts, v)
	}

	payRes, err := l.svcCtx.EmailRpc.OrderPaymentQueueEmail(l.ctx, &email.OrderPaymentQueueEmailRequest{OrderId: paymentOrderIds, Contact: paymentOrderContacts})
	if err != nil {
		logx.Alert("Send Daily unpaid order reminder email failed" + payRes.Msg)
		return err
	} else {
		logx.Info("Daily unpaid order reminder email sent.")
	}

	return nil
}
