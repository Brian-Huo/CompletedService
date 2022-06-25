package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCompanyLogic {
	return &ListCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCompanyLogic) ListCompany(req *types.ListCompanyRequest) (resp *types.ListCompanyResponse, err error) {
	// todo: add your logic here and delete this line

	return nil, status.Error(500, "Invalid, Currently not available")
}
