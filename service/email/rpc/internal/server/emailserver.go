// Code generated by goctl. DO NOT EDIT!
// Source: email.proto

package server

import (
	"context"

	"cleaningservice/service/email/rpc/internal/logic"
	"cleaningservice/service/email/rpc/internal/svc"
	"cleaningservice/service/email/rpc/types/email"
)

type EmailServer struct {
	svcCtx *svc.ServiceContext
	email.UnimplementedEmailServer
}

func NewEmailServer(svcCtx *svc.ServiceContext) *EmailServer {
	return &EmailServer{
		svcCtx: svcCtx,
	}
}

func (s *EmailServer) Announcement(ctx context.Context, in *email.AnnouncementRequest) (*email.AnnouncementResponse, error) {
	l := logic.NewAnnouncementLogic(ctx, s.svcCtx)
	return l.Announcement(in)
}

func (s *EmailServer) GeneralEmail(ctx context.Context, in *email.GeneralEmailRequest) (*email.GeneralEmailResponse, error) {
	l := logic.NewGeneralEmailLogic(ctx, s.svcCtx)
	return l.GeneralEmail(in)
}

func (s *EmailServer) InvoiceEmail(ctx context.Context, in *email.InvoiceEmailRequest) (*email.InvoiceEmailResponse, error) {
	l := logic.NewInvoiceEmailLogic(ctx, s.svcCtx)
	return l.InvoiceEmail(in)
}

func (s *EmailServer) OrderAwaitQueueEmail(ctx context.Context, in *email.OrderAwaitQueueEmailRequest) (*email.OrderAwaitQueueEmailResponse, error) {
	l := logic.NewOrderAwaitQueueEmailLogic(ctx, s.svcCtx)
	return l.OrderAwaitQueueEmail(in)
}

func (s *EmailServer) OrderPaymentQueueEmail(ctx context.Context, in *email.OrderPaymentQueueEmailRequest) (*email.OrderPaymentQueueEmailResponse, error) {
	l := logic.NewOrderPaymentQueueEmailLogic(ctx, s.svcCtx)
	return l.OrderPaymentQueueEmail(in)
}

func (s *EmailServer) OrderTransferQueueEmail(ctx context.Context, in *email.OrderTransferQueueEmailRequest) (*email.OrderTransferQueueEmailResponse, error) {
	l := logic.NewOrderTransferQueueEmailLogic(ctx, s.svcCtx)
	return l.OrderTransferQueueEmail(in)
}
