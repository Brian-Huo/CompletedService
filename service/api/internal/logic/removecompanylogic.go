package logic

import (
	"context"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/company"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemoveCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveCompanyLogic {
	return &RemoveCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveCompanyLogic) RemoveCompany(req *types.RemoveCompanyRequest) (resp *types.RemoveCompanyResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Company {
		return nil, status.Error(401, "Invalid, Unauthorised action.")
	}

	comp, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == company.ErrNotFound {
			return nil, status.Error(404, "Invalid, Company not found.")
		}
	}

	comp.CompanyStatus = int64(variables.Abolished)

	err = l.svcCtx.BCompanyModel.Update(l.ctx, comp)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.BEmployeeModel.ResignByCompany(l.ctx, uid)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return
}
