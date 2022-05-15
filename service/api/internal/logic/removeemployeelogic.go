package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemoveEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveEmployeeLogic {
	return &RemoveEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveEmployeeLogic) RemoveEmployee(req *types.RemoveEmployeeRequest) (resp *types.RemoveEmployeeResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	if role != variables.Company {
		return nil, status.Error(401, "Invalid, Unauthorised action.")
	}

	emp, err := l.svcCtx.BEmployeeModel.FindOne(l.ctx, req.Employee_id)
	if err != nil {
		if err == employee.ErrNotFound {
			return nil, status.Error(404, "Invalid, Employee not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if uid != emp.CompanyId {
		return nil, status.Error(404, "Invalid, Employee not found.")
	}

	emp.WorkStatus = int64(variables.Resigned)

	err = l.svcCtx.BEmployeeModel.Update(l.ctx, emp)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return
}
