package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEmployeeLogic {
	return &ListEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListEmployeeLogic) ListEmployee(req *types.ListEmployeeRequest) (resp *types.ListEmployeeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
