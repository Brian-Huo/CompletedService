package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/company"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DetailCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailCompanyLogic {
	return &DetailCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailCompanyLogic) DetailCompany(req *types.DetailCompanyRequest) (resp *types.DetailCompanyResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	res, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, req.Company_id)
	if err != nil {
		if err == company.ErrNotFound {
			return nil, status.Error(404, "Invalid, Company not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if role != variables.Company || uid != req.Company_id {
		res.PaymentId = sql.NullInt64{0, false}
		res.DepositeRate = -1
	}

	return &types.DetailCompanyResponse{
		Company_id:         res.CompanyId,
		Company_name:       res.CompanyName,
		Payment_id:         res.PaymentId.Int64,
		Director_name:      res.DirectorName.String,
		Contact_details:    res.ContactDetails,
		Registered_address: res.RegisteredAddress.Int64,
		Deposite_rate:      resp.Deposite_rate,
	}, nil
}
