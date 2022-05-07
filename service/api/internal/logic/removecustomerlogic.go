package logic

import (
	"context"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

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
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)

	if role != variables.Customer {
		return nil, status.Error(401, "Invalid, Unauthorised action.")
	}

	go l.svcCtx.RCustomerPaymentModel.DeleteByCustomer(l.ctx, uid)
	go l.svcCtx.RCustomerAddressModel.DeleteByCustomer(l.ctx, uid)
	// go l.svcCtx.BCustomerModel.Delete(l.ctx, uid)

	return
}
