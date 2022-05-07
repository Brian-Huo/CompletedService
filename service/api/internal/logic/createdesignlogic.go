package logic

import (
	"context"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/design"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDesignLogic {
	return &CreateDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDesignLogic) CreateDesign(req *types.CreateDesignRequest) (resp *types.CreateDesignResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Company {
		return nil, status.Error(401, "Invalid, Not company.")
	}

	_, err = l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, status.Error(404, "Invalid, Company not found.")
	}

	_, err = l.svcCtx.BServiceModel.FindOne(l.ctx, req.Service_id)
	if err != nil {
		return nil, status.Error(404, "Invalid, Service not found.")
	}

	newItem := design.BDesign{
		CompanyId: uid,
		ServiceId: req.Service_id,
		Price:     req.Price,
		Comments:  req.Comments,
	}

	res, err := l.svcCtx.BDesignModel.Insert(l.ctx, &newItem)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateDesignResponse{
		Design_id: newId,
	}, nil
}
