package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveContractorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveContractorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveContractorLogic {
	return &RemoveContractorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveContractorLogic) RemoveContractor(req *types.RemoveContractorRequest) (resp *types.RemoveContractorResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
