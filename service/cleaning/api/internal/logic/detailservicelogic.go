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

type DetailServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailServiceLogic {
	return &DetailServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailServiceLogic) DetailService(req *types.DetailServiceRequest) (service_itemp *types.DetailServiceResponse, err error) {
	service_item, err := l.svcCtx.BServiceModel.FindOne(l.ctx, req.Service_id)
	if err != nil {
		if err == service.ErrNotFound {
			return nil, status.Error(404, "Invalid, Service not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	category_item, err := l.svcCtx.BCategoryModel.FindOne(l.ctx, service_item.ServiceType)
	if err != nil {
		if err == category.ErrNotFound {
			return nil, status.Error(404, "Invalid, Category not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &types.DetailServiceResponse{
		Service_id: service_item.ServiceId,
		Service_type: types.DetailCategoryResponse{
			Category_id:          category_item.CategoryId,
			Category_name:        category_item.CategoryName,
			Category_description: category_item.CategoryDescription,
		},
		Service_scope:       service_item.ServiceScope,
		Service_name:        service_item.ServiceName,
		Service_description: service_item.ServiceDescription,
		Service_price:       service_item.ServicePrice,
	}, nil
}
