package logic

import (
	"context"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employeeservice"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateEmployeeServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEmployeeServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmployeeServiceLogic {
	return &CreateEmployeeServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEmployeeServiceLogic) CreateEmployeeService(req *types.CreateEmployeeServiceRequest) (resp *types.CreateEmployeeServiceResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Employee {
		return nil, status.Error(401, "Invalid, Not employee.")
	}

	_, err = l.svcCtx.BServiceModel.FindOne(l.ctx, req.Service_id)
	if err != nil {
		return nil, status.Error(404, "Invalid, Service not found.")
	}

	newItem := employeeservice.REmployeeService{
		EmployeeId: uid,
		ServiceId:  req.Service_id,
	}

	res, err := l.svcCtx.REmployeeServiceModel.Insert(l.ctx, &newItem)
	_ = res
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateEmployeeServiceResponse{}, nil
}
