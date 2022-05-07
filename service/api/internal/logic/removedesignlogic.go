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

type RemoveDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveDesignLogic {
	return &RemoveDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveDesignLogic) RemoveDesign(req *types.RemoveDesignRequest) (resp *types.RemoveDesignResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)

	if role != variables.Company {
		return nil, status.Error(401, "Invalid, Unauthoried action.")
	}

	des, err := l.svcCtx.BDesignModel.FindOne(l.ctx, req.Design_id)
	if err != nil {
		if err == design.ErrNotFound {
			return nil, status.Error(404, "Invalid, Design not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if uid != des.CompanyId {
		return nil, status.Error(404, "Invalid, Design not found.")
	}

	err = l.svcCtx.BDesignModel.Delete(l.ctx, req.Design_id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return
}
