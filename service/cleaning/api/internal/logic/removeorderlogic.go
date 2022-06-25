package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemoveOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveOrderLogic {
	return &RemoveOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveOrderLogic) RemoveOrder(req *types.RemoveOrderRequest) (resp *types.RemoveOrderResponse, err error) {
	// todo: add your logic here and delete this line

	return nil, status.Error(500, "Invalid, Currently not available")
}
