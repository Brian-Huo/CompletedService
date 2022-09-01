package logic

import (
	"context"
	"fmt"

	"cleaningservice/common/variables"
	"cleaningservice/service/email/rpc/internal/svc"
	"cleaningservice/service/email/rpc/types/email"
	"cleaningservice/util"

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
	// Email constant variables
	subject := "[Daily]Unpaid Orders Reminder - Daily Report"
	greetings := fmt.Sprintf("<p>Hi %s Reception Team,</p><br>", variables.Business_name)
	endings := "</br> Above is all unpaid orders currently within the system. Please be aware of their payment due dates and payment status. </br>"

	// Email main contents
	contents := ""

	for index, order_id := range in.OrderId {
		contents += fmt.Sprintf("<b>order</b> %s is due on %s day(s). If due date is near, please contact: %s <br>", order_id, in.DueDate[index], in.Contact[index])
	}

	contents += "</br> Please inform your manager and make sure these order(s) are fully reviewed.</br>"

	// Send email
	go util.SendToReception(subject, greetings+contents+endings)

	return &email.OrderPaymentQueueEmailResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
