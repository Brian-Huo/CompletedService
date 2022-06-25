package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemoveCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveCustomerLogic {
	return &RemoveCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveCustomerLogic) RemoveCustomer(req *types.RemoveCustomerRequest) (resp *types.RemoveCustomerResponse, err error) {
	return nil, status.Error(500, err.Error())
}
