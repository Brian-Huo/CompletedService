package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/design"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DetailDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailDesignLogic {
	return &DetailDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailDesignLogic) DetailDesign(req *types.DetailDesignRequest) (resp *types.DetailDesignResponse, err error) {
	res, err := l.svcCtx.BDesignModel.FindOne(l.ctx, req.Design_id)
	if err != nil {
		if err == design.ErrNotFound {
			return nil, status.Error(404, "Invalid, Design not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &types.DetailDesignResponse{
		Design_id:  res.DesignId,
		Service_id: res.ServiceId,
		Company_id: res.CompanyId,
		Price:      res.Price,
		Comments:   res.Comments,
	}, nil
}
