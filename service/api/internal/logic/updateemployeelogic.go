package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/cryptx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"
	"cleaningservice/service/model/employeeservice"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UpdateEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmployeeLogic {
	return &UpdateEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEmployeeLogic) UpdateEmployee(req *types.UpdateEmployeeRequest) (resp *types.UpdateEmployeeResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	var employeeId int64
	if role == variables.Employee {
		employeeId = uid
	} else if role == variables.Company {
		employeeId = req.Employee_id
	} else {
		return nil, status.Error(401, "Not Company/Employee.")
	}

	// Get employee details
	empl, err := l.svcCtx.BEmployeeModel.FindOne(l.ctx, employeeId)
	if err != nil {
		if err == employee.ErrNotFound {
			return nil, status.Error(404, "Invalid, Employee not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Verify company and employee
	if role == variables.Employee {
		empl.LinkCode = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Link_code)
	} else if role == variables.Company {
		if empl.CompanyId != uid {
			return nil, status.Error(404, "Invalid, Employee not found.")
		}
	}

	// Update employee details
	err = l.svcCtx.BEmployeeModel.Update(l.ctx, &employee.BEmployee{
		EmployeeId:     employeeId,
		EmployeePhoto:  sql.NullString{req.Employee_photo, req.Employee_photo != ""},
		EmployeeName:   req.Employee_name,
		ContactDetails: req.Contact_details,
		CompanyId:      empl.CompanyId,
		LinkCode:       empl.LinkCode,
		WorkStatus:     empl.WorkStatus,
		OrderId:        empl.OrderId,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Add new employee services
	for _, new_service := range req.New_services {
		_, err = l.svcCtx.REmployeeServiceModel.Insert(l.ctx, &employeeservice.REmployeeService{
			EmployeeId: employeeId,
			ServiceId:  new_service,
		})
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	}

	// Remove old employee services
	for _, old_service := range req.Remove_services {
		err = l.svcCtx.REmployeeServiceModel.Delete(l.ctx, employeeId, old_service)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	}

	return &types.UpdateEmployeeResponse{}, nil
}
