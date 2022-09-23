package logic

import (
	"context"
	"strings"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/category"
	"cleaningservice/service/cleaning/model/property"
	"cleaningservice/service/cleaning/model/region"
	"cleaningservice/service/cleaning/model/service"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type EnquireServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnquireServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnquireServiceLogic {
	return &EnquireServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnquireServiceLogic) EnquireService(req *types.EnquireServiceRequest) (resp *types.ListServiceResponse, err error) {
	// Get region details
	region_item, err := l.svcCtx.BRgionModel.FindOneByPostcode(l.ctx, req.Postcode)
	if err != nil {
		if err != region.ErrNotFound {
			return nil, status.Error(500, err.Error())
		}
		region_item = &region.BRegion{ChargeAmount: 0}
	}

	// Get property details
	property_item, err := l.svcCtx.BPorpertyModel.FindOneByPropertyName(l.ctx, strings.ToLower(req.Property))
	if err != nil {
		if err == property.ErrNotFound {
			return nil, status.Error(404, "Invalid, Property type not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Get category details
	category_item, err := l.svcCtx.BCategoryModel.FindOne(l.ctx, req.Category_id)
	if err != nil {
		if err == category.ErrNotFound {
			return nil, status.Error(404, "Invalid, Category not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// List all service in this category
	service_items, err := l.svcCtx.BServiceModel.FindAllByCategory(l.ctx, category_item.CategoryId)
	if err != nil {
		if err == service.ErrNotFound {
			return nil, status.Error(404, "Invalid, Service not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailServiceResponse{}

	for _, item := range service_items {
		service_response := types.DetailServiceResponse{
			Service_id: item.ServiceId,
			Service_type: types.DetailCategoryResponse{
				Category_id:          category_item.CategoryId,
				Category_addr:        category_item.CategoryAddr,
				Category_name:        category_item.CategoryName,
				Category_description: category_item.CategoryDescription,
			},
			Service_scope:       item.ServiceScope,
			Service_name:        item.ServiceName,
			Service_photo:       item.ServicePhoto.String,
			Service_description: item.ServiceDescription,
			Service_price:       item.ServicePrice * float64(1+(region_item.ChargeAmount+property_item.ChargeAmount)/100),
		}

		allItems = append(allItems, service_response)
	}

	return &types.ListServiceResponse{
		Items: allItems,
	}, nil
}
