package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/category"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DetailCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailCategoryLogic {
	return &DetailCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailCategoryLogic) DetailCategory(req *types.DetailCategoryRequest) (resp *types.DetailCategoryResponse, err error) {
	category_item, err := l.svcCtx.BCategoryModel.FindOne(l.ctx, req.Category_id)
	if err != nil {
		if err == category.ErrNotFound {
			return nil, status.Error(404, "Invalid, Category not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &types.DetailCategoryResponse{
		Category_id:          category_item.CategoryId,
		Category_addr:        category_item.CategoryAddr,
		Category_name:        category_item.CategoryName,
		Category_description: category_item.CategoryDescription,
	}, nil
}
