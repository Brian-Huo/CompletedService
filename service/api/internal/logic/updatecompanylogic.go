package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/company"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompanyLogic {
	return &UpdateCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCompanyLogic) UpdateCompany(req *types.UpdateCompanyRequest) (resp *types.UpdateCompanyResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	} else if role != variables.Company {
		return nil, errorx.NewCodeError(401, "Invalid, Not company.")
	}

	company_item, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == company.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Company not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	company_item.CompanyName = req.Company_name
	company_item.DirectorName = sql.NullString{req.Director_name, req.Director_name != ""}
	company_item.ContactDetails = req.Contact_details

	err = l.svcCtx.BCompanyModel.Update(l.ctx, company_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	return &types.UpdateCompanyResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
