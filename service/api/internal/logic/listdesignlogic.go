package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/design"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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
	res, err := l.svcCtx.BDesignModel.List(l.ctx)
	if err != nil {
		if err == design.ErrNotFound {
			return nil, status.Error(404, "Invalid, Design not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailDesignResponse{}

	for _, item := range res {
		newItem := types.DetailDesignResponse{
			Design_id:  item.DesignId,
			Company_id: item.CompanyId,
			Service_id: item.ServiceId,
			Price:      item.Price,
			Comments:   item.Comments,
		}

		allItems = append(allItems, newItem)
	}

	return &types.ListDesignResponse{
		Items: allItems,
	}, nil
}
