package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListPaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPaymentLogic {
	return &ListPaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPaymentLogic) ListPayment(req *types.ListPaymentRequest) (resp *types.ListPaymentResponse, err error) {
	return nil, status.Error(500, err.Error())
}
