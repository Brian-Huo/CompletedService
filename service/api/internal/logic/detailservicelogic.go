package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/service"

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

func (l *DetailServiceLogic) DetailService(req *types.DetailServiceRequest) (resp *types.DetailServiceResponse, err error) {
	res, err := l.svcCtx.BServiceModel.FindOne(l.ctx, req.Service_id)
	if err != nil {
		if err == service.ErrNotFound {
			return nil, status.Error(404, "Invalid, Service not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &types.DetailServiceResponse{
		Service_id:          res.ServiceId,
		Service_type:        res.ServiceType,
		Service_scope:       res.ServiceScope,
		Service_name:        res.ServiceName,
		Service_photo:       res.ServicePhoto,
		Service_description: res.ServiceDescription,
		Service_price:       res.ServicePrice,
	}, nil
}
