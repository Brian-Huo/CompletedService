package logic

import (
	"context"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/company"

	"github.com/zeromicro/go-zero/core/logx"
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
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	} else if role != variables.Company {
		return nil, errorx.NewCodeError(401, "Invalid, Unauthorised action.")
	}

	company_item, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == company.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Company not found.")
		}
	}

	company_item.FinanceStatus = company.Abolished

	err = l.svcCtx.BCompanyModel.Update(l.ctx, company_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	err = l.svcCtx.BContractorModel.ResignByFinance(l.ctx, uid)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	return &types.RemoveCompanyResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
