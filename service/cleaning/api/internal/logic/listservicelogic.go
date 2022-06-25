package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/category"
	"cleaningservice/service/cleaning/model/service"

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

func (l *ListServiceLogic) ListService(req *types.ListServiceRequest) (service_itemp *types.ListServiceResponse, err error) {
	service_items, err := l.svcCtx.BServiceModel.List(l.ctx)
	if err != nil {
		if err == service.ErrNotFound {
			return nil, status.Error(404, "Invalid, Service not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailServiceResponse{}

	for _, item := range service_items {
		category_item, err := l.svcCtx.BCategoryModel.FindOne(l.ctx, item.ServiceType)
		if err != nil {
			if err == category.ErrNotFound {
				return nil, status.Error(404, "Invalid, Category not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		service_response := types.DetailServiceResponse{
			Service_id: item.ServiceId,
			Service_type: types.DetailCategoryResponse{
				Category_id:          category_item.CategoryId,
				Category_name:        category_item.CategoryName,
				Category_description: category_item.CategoryDescription,
			},
			Service_scope:       item.ServiceScope,
			Service_name:        item.ServiceName,
			Service_photo:       item.ServicePhoto.String,
			Service_description: item.ServiceDescription,
			Service_price:       item.ServicePrice,
		}

		allItems = append(allItems, service_response)
	}

	return &types.ListServiceResponse{
		Items: allItems,
	}, nil
}
