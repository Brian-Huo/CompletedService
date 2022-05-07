package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListServiceLogic {
	return &ListServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListServiceLogic) ListService(req *types.ListServiceRequest) (resp *types.ListServiceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
