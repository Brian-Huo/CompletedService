package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDesignLogic {
	return &ListDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDesignLogic) ListDesign(req *types.ListDesignRequest) (resp *types.ListDesignResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
