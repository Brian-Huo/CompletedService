package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/company"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type LoginServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginServiceLogic {
	return &LoginServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginServiceLogic) LoginService(req *types.LoginServiceRequest) (resp *types.LoginServiceResponse, err error) {
	// find company by contact_details
	item, err := l.svcCtx.BCompanyModel.FindOnebyPhone(l.ctx, req.Contact_details)
	if err == company.ErrNotFound {
		// if company not found, insert company
		res, err := l.svcCtx.BCompanyModel.Insert(l.ctx, &company.BCompany{
			CompanyName:       req.Contact_details,
			PaymentId:         sql.NullInt64{0, false},
			DirectorName:      sql.NullString{"", false},
			ContactDetails:    req.Contact_details,
			RegisteredAddress: sql.NullInt64{0, false},
			DepositeRate:      10,
		})
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		// get company id
		newId, err := res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		item.CompanyId = newId
	} else if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 签发 jwt token
	token, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, 1, l.svcCtx.Config.Auth.AccessExpire,
		item.CompanyId, variables.Company)
	if err != nil {
		return nil, status.Error(500, "Jwt token error.")
	}

	return &types.LoginServiceResponse{
		Code:        "200",
		Message:     "success",
		AccessToken: token,
	}, nil
}
