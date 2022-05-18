package logic

import (
	"context"
	"fmt"
	"time"

	"cleaningservice/common/cryptx"
	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginEmployeeLogic {
	return &LoginEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginEmployeeLogic) LoginEmployee(req *types.LoginEmployeeRequest) (resp *types.LoginEmployeeResponse, err error) {
	// find employee by contact_details
	item, err := l.svcCtx.BEmployeeModel.FindOnebyPhone(l.ctx, req.Contact_details)
	if err != nil {
		if err == employee.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Employee not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	fmt.Println(len(item.LinkCode), "      ", len(req.LinkCode))
	// Verify if first login
	if item.WorkStatus == int64(variables.Await) {
		if req.VerifyCode == "" {
			return nil, errorx.NewCodeError(1002, "Invalid, Verify code required.")
		} else if false {
			return nil, errorx.NewCodeError(401, "Invalid, Verfiy code incorrect.")
		}

		if item.LinkCode != req.LinkCode {
			return nil, errorx.NewCodeError(401, "Invalid, Link code incorrect.")
		} else {
			// Update employee first update
			item.WorkStatus = int64(variables.Vacant)
			item.LinkCode = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, item.LinkCode)
			err = l.svcCtx.BEmployeeModel.Update(l.ctx, item)
			if err != nil {
				return nil, errorx.NewCodeError(500, err.Error())
			}
		}
	} else if item.LinkCode != cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.LinkCode) {
		return nil, errorx.NewCodeError(401, "Invalid, Link code incorrect.")
	}

	// 签发 jwt token
	now := time.Now().Unix()
	token, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire,
		item.EmployeeId, variables.Contractor)
	if err != nil {
		return nil, errorx.NewCodeError(500, "Jwt token error.")
	}

	return &types.LoginEmployeeResponse{
		Code:        200,
		Message:     "success",
		AccessToken: token,
	}, nil
}
