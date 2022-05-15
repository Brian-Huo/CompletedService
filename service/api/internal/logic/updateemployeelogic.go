package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"

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
	} else if role == variables.Customer {
		return nil, status.Error(401, "Not Company/Employee.")
	} else if role == variables.Employee {
		empl, err := l.svcCtx.BEmployeeModel.FindOne(l.ctx, uid)
		if err != nil {
			if err == employee.ErrNotFound {
				return nil, status.Error(404, "Invalid, Employee not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		err = l.svcCtx.BEmployeeModel.Update(l.ctx, &employee.BEmployee{
			EmployeeId:     uid,
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
	} else if role == variables.Company {
		empl, err := l.svcCtx.BEmployeeModel.FindOne(l.ctx, req.Employee_id)
		if err != nil {
			if err == employee.ErrNotFound {
				return nil, status.Error(404, "Invalid, Employee not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		if empl.CompanyId != uid {
			return nil, status.Error(404, "Invalid, Employee not found.")
		}

		err = l.svcCtx.BEmployeeModel.Update(l.ctx, &employee.BEmployee{
			EmployeeId:     req.Employee_id,
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
	}

	return &types.UpdateEmployeeResponse{}, nil
}
