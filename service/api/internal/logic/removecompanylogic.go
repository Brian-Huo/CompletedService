package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemoveCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveCompanyLogic {
	return &RemoveCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveCompanyLogic) RemoveCompany(req *types.RemoveCompanyRequest) (resp *types.RemoveCompanyResponse, err error) {
	// todo: add your logic here and delete this line

	return nil, status.Error(500, "Invalid, Currently not available")
}
