package logic

import (
	"context"
	"time"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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
			return nil, status.Error(404, err.Error())
		}
		return nil, status.Error(500, err.Error())
	}

	// 签发 jwt token
	now := time.Now().Unix()
	token, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire,
		item.EmployeeId, variables.Company)
	if err != nil {
		return nil, status.Error(500, "Jwt token error.")
	}

	return &types.LoginEmployeeResponse{
		Code:        "200",
		Message:     "success",
		AccessToken: token,
	}, nil
}
