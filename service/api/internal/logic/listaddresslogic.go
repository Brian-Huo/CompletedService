package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAddressLogic {
	return &ListAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAddressLogic) ListAddress(req *types.ListAddressRequest) (resp *types.ListAddressResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
