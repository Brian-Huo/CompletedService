package logic

import (
	"context"
	"database/sql"
	"time"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/company"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type LoginCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginCompanyLogic {
	return &LoginCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginCompanyLogic) LoginCompany(req *types.LoginCompanyRequest) (resp *types.LoginCompanyResponse, err error) {
	// find company by contact_details
	var companyId int64
	res, err := l.svcCtx.BCompanyModel.FindOneByContactDetails(l.ctx, req.Contact_details)
	if err == company.ErrNotFound {
		// Passpath check
		if req.VerifyCode != "Ae69Ziud" {
			return nil, status.Error(500, err.Error())
		}

		// if company not found, insert company
		company_item, err := l.svcCtx.BCompanyModel.Insert(l.ctx, &company.BCompany{
			CompanyName:       req.Contact_details,
			PaymentId:         sql.NullInt64{Int64: 0, Valid: false},
			DirectorName:      sql.NullString{String: "", Valid: false},
			ContactDetails:    req.Contact_details,
			RegisteredAddress: 1,
			DepositeRate:      10,
			FinanceStatus:     company.Active,
		})
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		// get company id
		newId, err := company_item.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		companyId = newId
	} else if err == nil {
		companyId = res.CompanyId
	} else {
		return nil, status.Error(500, err.Error())
	}

	// 签发 jwt token
	now := time.Now().Unix()
	token, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire,
		companyId, variables.Company)
	if err != nil {
		return nil, status.Error(500, "Jwt token error.")
	}

	return &types.LoginCompanyResponse{
		Code:        200,
		Msg:         "success",
		AccessToken: token,
	}, nil
}
