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
	// Email constant variables
	subject := "[Urgent]Order Transfer Reminder - Order ID: " + in.OrderId
	greetings := fmt.Sprintf("<p>Hi %s Reception Team,</p><br>", variables.Business_name)
	endings := "</br> Please inform your manager and negotiate with our customers ASAP.</br>"

	// Email main contents
	contents := fmt.Sprintf("<b>order</b> %s is requiring immediately transfer with contact details: %s.<br>", in.OrderId, in.Contact)

	// Send email
	go util.SendToReception(subject, greetings+contents+endings)

	return &email.OrderTransferQueueEmailResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
