package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/category"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoryLogic {
	return &ListCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCategoryLogic) ListCategory(req *types.ListCategoryRequest) (resp *types.ListCategoryResponse, err error) {
	category_items, err := l.svcCtx.BCategoryModel.List(l.ctx)
	if err != nil {
		if err == category.ErrNotFound {
			return nil, status.Error(404, "Invalid, Category not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailCategoryResponse{}

	for _, item := range category_items {
		newItem := types.DetailCategoryResponse{
			Category_id:   item.CategoryId,
			Category_name: item.CategoryName,
		}

		allItems = append(allItems, newItem)
	}

	return &types.ListCategoryResponse{
		Items: allItems,
	}, nil
}
