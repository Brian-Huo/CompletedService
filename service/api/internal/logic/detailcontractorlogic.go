package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailContractorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailContractorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailContractorLogic {
	return &DetailContractorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailContractorLogic) DetailContractor(req *types.DetailContractorRequest) (resp *types.DetailContractorResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
