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
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Company {
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
		newItem := types.DetailEmployeeResponse{
			Employee_id:     item.EmployeeId,
			Employee_photo:  item.EmployeePhoto.String,
			Employee_name:   item.EmployeeName,
			Contact_details: item.ContactDetails,
			Company_id:      item.CompanyId,
			Link_code:       item.LinkCode,
			Work_status:     int(item.WorkStatus),
			Order_id:        item.OrderId.Int64,
		}

		allItems = append(allItems, newItem)
	}

	return &types.ListEmployeeResponse{
		Items: allItems,
	}, nil
}
