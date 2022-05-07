package logic

import (
	"context"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"
	"cleaningservice/service/model/employeeservice"
	"cleaningservice/service/model/service"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListEmployeeServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListEmployeeServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEmployeeServiceLogic {
	return &ListEmployeeServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListEmployeeServiceLogic) ListEmployeeService(req *types.ListEmployeeServiceRequest) (resp *types.ListEmployeeServiceResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Company {
		return nil, status.Error(401, "Invalid, Not company.")
	}

	empl, err := l.svcCtx.BEmployeeModel.FindOne(l.ctx, req.Employee_id)
	if err != nil {
		if err == employee.ErrNotFound {
			return nil, status.Error(404, "Invalid, Employee not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if uid != empl.CompanyId {
		return nil, status.Error(404, "Invalid, Employee not found.")
	}

	res, err := l.svcCtx.REmployeeServiceModel.FindAllByEmployee(l.ctx, req.Employee_id)
	if err != nil {
		if err == employeeservice.ErrNotFound {
			return nil, status.Error(404, "Invalid, Employee service not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailServiceResponse{}

	for _, item := range res {
		serv, err := l.svcCtx.BServiceModel.FindOne(l.ctx, item.ServiceId)
		if err != nil {
			if err == service.ErrNotFound {
				return nil, status.Error(404, "Invalid, Service not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		newItem := types.DetailServiceResponse{
			Service_id:          serv.ServiceId,
			Service_type:        serv.ServiceType,
			Service_description: serv.ServiceDescription.String,
		}

		allItems = append(allItems, newItem)
	}

	return &types.ListEmployeeServiceResponse{
		Items: allItems,
	}, nil

	return
}
