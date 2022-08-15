package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContractorHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContractorHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContractorHistoryLogic {
	return &GetContractorHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContractorHistoryLogic) GetContractorHistory(req *types.GetContractorHistoryRequest) (resp *types.GetContractorHistoryResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
