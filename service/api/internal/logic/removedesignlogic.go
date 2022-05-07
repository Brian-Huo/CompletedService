package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveDesignLogic {
	return &RemoveDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveDesignLogic) RemoveDesign(req *types.RemoveDesignRequest) (resp *types.RemoveDesignResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
