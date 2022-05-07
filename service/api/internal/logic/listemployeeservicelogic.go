package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListEmployeeServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListEmployeeServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEmployeeServiceLogic {
	return &ListEmployeeServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListEmployeeServiceLogic) ListEmployeeService(req *types.ListEmployeeServiceRequest) (resp *types.ListEmployeeServiceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
