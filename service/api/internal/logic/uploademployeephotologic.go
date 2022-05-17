package logic

import (
	"context"
	"database/sql"
	"strconv"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"
	"cleaningservice/util"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UploadEmployeePhotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadEmployeePhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadEmployeePhotoLogic {
	return &UploadEmployeePhotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadEmployeePhotoLogic) UploadEmployeePhoto(req *types.UploadEmployeePhotoRequest) (resp *types.UploadEmployeePhotoResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	var employeeId int64
	if role == variables.Company {
		employeeId = req.Employee_id
	} else if role == variables.Employee {
		employeeId = uid
	}

	// Get employee details
	empl, err := l.svcCtx.BEmployeeModel.FindOne(l.ctx, employeeId)
	if err != nil {
		if err == employee.ErrNotFound {
			return nil, status.Error(404, "Invalid, Employee not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Verify company
	if role == variables.Company {
		if empl.CompanyId != uid {
			return nil, status.Error(404, "Invalid, Employee not found.")
		}
	}

	// Save image
	photoPath, err := util.SaveImage(req.Employee_photo, "employee", strconv.FormatInt(employeeId, 10))
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Update employee details
	empl.EmployeePhoto = sql.NullString{photoPath, photoPath != strconv.FormatInt(employeeId, 10)}
	err = l.svcCtx.BEmployeeModel.Update(l.ctx, empl)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return
}
