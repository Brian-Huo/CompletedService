package logic

import (
	"context"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemoveEmployeeServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveEmployeeServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveEmployeeServiceLogic {
	return &RemoveEmployeeServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveEmployeeServiceLogic) RemoveEmployeeService(req *types.RemoveEmployeeServiceRequest) (resp *types.RemoveEmployeeServiceResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)

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

	err = l.svcCtx.REmployeeServiceModel.Delete(l.ctx, req.Employee_id, req.Service_id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return
}
