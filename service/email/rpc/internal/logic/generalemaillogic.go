package logic

import (
	"context"
	"log"

	"cleaningservice/common/variables"
	"cleaningservice/service/email/rpc/internal/svc"
	"cleaningservice/service/email/rpc/types/email"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/gomail.v2"
)

type GeneralEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const emailEnd = "<p>If you have any questions and concerns, you can kindly reply to this email</p><br><p>Kind regards,</p><br>"
const emailSign = variables.Business_name + " Support Team"

func NewGeneralEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeneralEmailLogic {
	return &GeneralEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GeneralEmailLogic) GeneralEmail(in *email.GeneralEmailRequest) (*email.GeneralEmailResponse, error) {
	// Set attributes
	m := gomail.NewMessage()
	m.SetHeader("From", variables.QME_email)
	m.SetHeader("To", in.Target)
	m.SetHeader("Subject", in.Subject)
	m.SetBody("text/html", in.Content+emailEnd+emailSign)

	// Send the email
	d := gomail.NewDialer("smtp.gmail.com", 587, variables.QME_email, variables.QME_password)
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
		return &email.GeneralEmailResponse{
			Code: 500,
			Msg:  "Send general email failed",
		}, err
	}

	return &email.GeneralEmailResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
