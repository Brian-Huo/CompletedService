package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/company"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Company {
		return nil, status.Error(401, "Invalid, Not company.")
	}

	res, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, req.Company_id)
	if err != nil {
		if err == company.ErrNotFound {
			return nil, status.Error(404, "Company not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.BCompanyModel.Update(l.ctx, &company.BCompany{
		CompanyId:         uid,
		CompanyName:       req.Company_name,
		PaymentId:         sql.NullInt64{req.Payment_id, req.Payment_id != 0},
		DirectorName:      sql.NullString{req.Director_name, req.Director_name != ""},
		ContactDetails:    req.Contact_details,
		RegisteredAddress: sql.NullInt64{req.Registered_address, req.Registered_address != 0},
		DepositeRate:      res.DepositeRate,
	})

	if err != nil {
		if err == company.ErrNotFound {
			return nil, status.Error(404, "Invalid, Company not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &types.UpdateCompanyResponse{}, nil
}
