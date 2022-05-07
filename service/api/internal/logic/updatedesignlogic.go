package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/design"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UpdateDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDesignLogic {
	return &UpdateDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDesignLogic) UpdateDesign(req *types.UpdateDesignRequest) (resp *types.UpdateDesignResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Company {
		return nil, status.Error(401, "Invalid, Not company.")
	}

	res, err := l.svcCtx.BDesignModel.FindOne(l.ctx, req.Design_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(404, "Design not found.")
		}
		return nil, status.Error(500, err.Error())
	}
	if res.CompanyId != uid {
		return nil, status.Error(401, "Invalid company design id.")
	}

	err = l.svcCtx.BDesignModel.Update(l.ctx, &design.BDesign{
		DesignId:  req.Design_id,
		CompanyId: req.Company_id,
		ServiceId: req.Service_id,
		Price:     req.Price,
		Comments:  req.Comments,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.UpdateDesignResponse{}, nil
}
