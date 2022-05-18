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

type DetailEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailEmployeeLogic {
	return &DetailEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailEmployeeLogic) DetailEmployee(req *types.DetailEmployeeRequest) (resp *types.DetailEmployeeResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	var res *employee.BEmployee
	if role != variables.Employee {
		res, err = l.svcCtx.BEmployeeModel.FindOne(l.ctx, req.Employee_id)
		if err != nil {
			if err == employee.ErrNotFound {
				return nil, status.Error(404, "Invalid, Employee not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		if role == variables.Customer {
			res.LinkCode = ""
			res.OrderId = sql.NullInt64{0, false}
			res.WorkStatus = -1
		}
	} else {
		res, err = l.svcCtx.BEmployeeModel.FindOne(l.ctx, uid)
		if err != nil {
			if err == employee.ErrNotFound {
				return nil, status.Error(404, "Invalid, Employee not found.")
			}
			return nil, status.Error(500, err.Error())
		}
	}

	// Get all emplooyee service
	service_list := types.ListEmployeeServiceResponse{}
	service_res, err := l.svcCtx.REmployeeServiceModel.FindAllByEmployee(l.ctx, res.EmployeeId)
	if err == nil {
		allServices := []types.DetailServiceResponse{}
		for _, res_item := range service_res {
			service_item, err := l.svcCtx.BServiceModel.FindOne(l.ctx, res_item.ServiceId)
			if err != nil {
				break
			}

			service := types.DetailServiceResponse{
				Service_id:          service_item.ServiceId,
				Service_type:        service_item.ServiceType,
				Service_description: service_item.ServiceDescription,
				Service_price:       service_item.ServicePrice,
			}

			allServices = append(allServices, service)
		}
		service_list.Items = allServices
	}

	return &types.DetailEmployeeResponse{
		Employee_id:      res.EmployeeId,
		Employee_photo:   res.EmployeePhoto.String,
		Employee_name:    res.EmployeeName,
		Contact_details:  res.ContactDetails,
		Company_id:       res.CompanyId,
		Link_code:        res.LinkCode,
		Work_status:      int(res.WorkStatus),
		Order_id:         res.OrderId.Int64,
		Employee_service: service_list,
	}, nil
}
