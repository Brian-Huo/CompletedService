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

type ListEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEmployeeLogic {
	return &ListEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListEmployeeLogic) ListEmployee(req *types.ListEmployeeRequest) (resp *types.ListEmployeeResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, Json format error")
	} else if role != variables.Company {
		return nil, status.Error(401, "Invalid, Not company.")
	}

	res, err := l.svcCtx.BEmployeeModel.FindAllByCompany(l.ctx, uid)
	if err != nil {
		if err == employee.ErrNotFound {
			return nil, status.Error(404, "Invalid, Employee not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailEmployeeResponse{}

	for _, item := range res {
		// Check if employee resigned
		if item.WorkStatus == int64(variables.Resigned) {
			continue
		}

		// Get all emplooyee service
		service_list := types.ListEmployeeServiceResponse{}
		service_res, err := l.svcCtx.REmployeeServiceModel.FindAllByEmployee(l.ctx, item.EmployeeId)
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

		// Get employee details
		newItem := types.DetailEmployeeResponse{
			Employee_id:      item.EmployeeId,
			Employee_photo:   item.EmployeePhoto.String,
			Employee_name:    item.EmployeeName,
			Contact_details:  item.ContactDetails,
			Company_id:       item.CompanyId,
			Link_code:        item.LinkCode,
			Work_status:      int(item.WorkStatus),
			Order_id:         item.OrderId.Int64,
			Employee_service: service_list,
		}

		allItems = append(allItems, newItem)
	}

	return &types.ListEmployeeResponse{
		Items: allItems,
	}, nil
}
