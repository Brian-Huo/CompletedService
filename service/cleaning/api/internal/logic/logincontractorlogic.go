package logic

import (
	"context"
	"fmt"
	"time"

	"cleaningservice/common/cryptx"
	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/contractor"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginContractorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginContractorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginContractorLogic {
	return &LoginContractorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginContractorLogic) LoginContractor(req *types.LoginContractorRequest) (resp *types.LoginContractorResponse, err error) {
	// find contractor by contact_details
	res, err := l.svcCtx.BContractorModel.FindOneByContactDetails(l.ctx, req.Contact_details)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Contractor not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	fmt.Println(len(res.LinkCode), "      ", len(req.LinkCode))
	// Verify if first login
	if res.WorkStatus == contractor.Await {
		if req.VerifyCode == "" {
			return nil, errorx.NewCodeError(1002, "Invalid, Verify code required.")
		} else if req.VerifyCode != "556" {
			return nil, errorx.NewCodeError(401, "Invalid, Verfiy code incorrect.")
		}

		if res.LinkCode != req.LinkCode {
			return nil, errorx.NewCodeError(401, "Invalid, Link code incorrect.")
		} else {
			// Update contractor first update
			res.WorkStatus = contractor.Vacant
			res.LinkCode = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, res.LinkCode)
			err = l.svcCtx.BContractorModel.Update(l.ctx, res)
			if err != nil {
				return nil, errorx.NewCodeError(500, err.Error())
			}
		}
	} else if res.WorkStatus == contractor.Resigned {
		return nil, errorx.NewCodeError(404, "Invalid, Contractor Not Found.")

	} else if res.LinkCode != cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.LinkCode) {
		return nil, errorx.NewCodeError(401, "Invalid, Link code incorrect.")
	}

	// 签发 jwt token
	now := time.Now().Unix()
	token, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire,
		res.ContractorId, variables.Contractor)
	if err != nil {
		return nil, errorx.NewCodeError(500, "Jwt token error.")
	}

	return &types.LoginContractorResponse{
		Code:        200,
		Msg:         "success",
		AccessToken: token,
	}, nil
}
