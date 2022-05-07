package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDesignLogic {
	return &UpdateDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDesignLogic) UpdateDesign(req *types.UpdateDesignRequest) (resp *types.UpdateDesignResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
