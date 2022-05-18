package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListContractorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListContractorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListContractorLogic {
	return &ListContractorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListContractorLogic) ListContractor(req *types.ListContractorRequest) (resp *types.ListContractorResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
