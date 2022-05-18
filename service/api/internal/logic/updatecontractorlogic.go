package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateContractorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateContractorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContractorLogic {
	return &UpdateContractorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateContractorLogic) UpdateContractor(req *types.UpdateContractorRequest) (resp *types.UpdateContractorResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
