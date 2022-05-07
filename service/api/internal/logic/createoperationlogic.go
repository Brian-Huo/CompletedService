package logic

import (
	"context"
	"time"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/operation"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateOperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOperationLogic {
	return &CreateOperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOperationLogic) CreateOperation(req *types.CreateOperationRequest) (resp *types.CreateOperationResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Employee {
		return nil, status.Error(401, "Invalid, Not employee.")
	}

	_, err = l.svcCtx.BEmployeeModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, status.Error(404, "Invalid, Employee not found.")
	}

	_, err = l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		return nil, status.Error(404, "Invalid, Order not found.")
	}

	issue_date, err := time.Parse("2006-01-02 15:04:05", req.Issue_date)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newItem := operation.BOperation{
		EmployeeId: uid,
		OrderId:    req.Order_id,
		Operation:  req.Operation,
		IssueDate:  issue_date,
	}

	res, err := l.svcCtx.BOperationModel.Insert(l.ctx, &newItem)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateOperationResponse{
		Operation_id: newId,
	}, nil

}
