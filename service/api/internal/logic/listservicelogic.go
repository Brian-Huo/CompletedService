package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/service"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListServiceLogic {
	return &ListServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListServiceLogic) ListService(req *types.ListServiceRequest) (resp *types.ListServiceResponse, err error) {
	res, err := l.svcCtx.BServiceModel.List(l.ctx)
	if err != nil {
		if err == service.ErrNotFound {
			return nil, status.Error(404, "Invalid, Service not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailServiceResponse{}

	for _, item := range res {
		newItem := types.DetailServiceResponse{
			Service_id:          item.ServiceId,
			Service_type:        item.ServiceType,
			Service_scope:       item.ServiceScope,
			Service_name:        item.ServiceName,
			Service_photo:       item.ServicePhoto,
			Service_description: item.ServiceDescription,
			Service_price:       item.ServicePrice,
		}

		allItems = append(allItems, newItem)
	}

	return &types.ListServiceResponse{
		Items: allItems,
	}, nil
}
