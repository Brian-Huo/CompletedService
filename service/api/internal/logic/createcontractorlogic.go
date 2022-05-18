package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContractorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContractorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContractorLogic {
	return &CreateContractorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContractorLogic) CreateContractor(req *types.CreateContractorRequest) (resp *types.CreateContractorResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
