package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"
	"cleaningservice/util"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmployeeLogic {
	return &CreateEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEmployeeLogic) CreateEmployee(req *types.CreateEmployeeRequest) (resp *types.CreateEmployeeResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Company {
		return nil, status.Error(401, "Invalid, Not company.")
	}

	_, err = l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, status.Error(404, "Invalid, Company not found.")
	}

	newItem := employee.BEmployee{
		EmployeePhoto:  sql.NullString{req.Employee_photo, req.Employee_photo != ""},
		EmployeeName:   req.Employee_name,
		ContactDetails: req.Contact_details,
		CompanyId:      uid,
		LinkCode:       util.RandStringBytesMaskImprSrcUnsafe(8),
		WorkStatus:     int64(variables.Vacant),
		OrderId:        sql.NullInt64{0, false},
	}

	res, err := l.svcCtx.BEmployeeModel.Insert(l.ctx, &newItem)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateEmployeeResponse{
		Employee_id: newId,
	}, nil
}
