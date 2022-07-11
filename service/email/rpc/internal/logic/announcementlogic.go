package logic

import (
	"context"

	"cleaningservice/service/email/rpc/internal/svc"
	"cleaningservice/service/email/rpc/types/email"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnnouncementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnnouncementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnnouncementLogic {
	return &AnnouncementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AnnouncementLogic) Announcement(in *email.AnnouncementRequest) (*email.AnnouncementResponse, error) {
	// todo: add your logic here and delete this line

	return &email.AnnouncementResponse{}, nil
}
