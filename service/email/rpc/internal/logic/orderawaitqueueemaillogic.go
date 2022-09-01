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
	// Email constant variables
	subject := "[Daily]In Queue Orders Reminder - Daily Report"
	greetings := fmt.Sprintf("<p>Hi %s Reception Team,</p><br>", variables.Business_name)
	endings := "</br> Above is all orders currently within the system and waiting for conttractors. Please be aware of their reservation, avoiding overdue and long staies. </br>"

	// Email main contents
	contents := ""

	for index, order_id := range in.OrderId {
		contents += fmt.Sprintf("<b>order</b> %s is remaining in the system for %s day(s). <br>", order_id, in.Vacancy[index])
	}

	contents += "</br> Please inform your manager and make sure these order(s) are fully reviewed.</br>"

	// Send email
	go util.SendToReception(subject, greetings+contents+endings)

	return &email.OrderAwaitQueueEmailResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
