package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/category"

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
		Category_id:   category_item.CategoryId,
		Category_name: category_item.CategoryName,
	}, nil
}
