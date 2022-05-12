package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemoveAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAddressLogic {
	return &RemoveAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveAddressLogic) RemoveAddress(req *types.RemoveAddressRequest) (resp *types.RemoveAddressResponse, err error) {
	return nil, status.Error(500, err.Error())
}
