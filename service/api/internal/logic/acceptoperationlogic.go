package logic

import (
	"context"
	"database/sql"
	"time"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"
	"cleaningservice/service/model/operation"
	"cleaningservice/service/model/order"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type AcceptOperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAcceptOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AcceptOperationLogic {
	return &AcceptOperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AcceptOperationLogic) AcceptOperation(req *types.AcceptOperationRequest) (resp *types.AcceptOperationResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Employee {
		return nil, status.Error(401, "Invalid, Not employee.")
	}

	// Check employee status
	empl, err := l.svcCtx.BEmployeeModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == employee.ErrNotFound {
			return nil, status.Error(404, "Invalid, Employee not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if empl.WorkStatus != int64(variables.Vacant) {
		return nil, status.Error(401, "Invalid, Employee should not double accept order(s).")
	}

	// Check order status
	ord, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if ord.Status != int64(variables.Queuing) {
		return nil, status.Error(401, "Order is currently unavailable.")
	}

	// Create operaction records
	newItem := operation.BOperation{
		EmployeeId: uid,
		OrderId:    req.Order_id,
		Operation:  int64(variables.Accept),
		IssueDate:  time.Now(),
	}

	res, err := l.svcCtx.BOperationModel.Insert(l.ctx, &newItem)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Update order details
	ord.EmployeeId = sql.NullInt64{uid, true}
	err = l.svcCtx.BOrderModel.Update(l.ctx, ord)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Update employee details
	empl.WorkStatus = int64(variables.InWork)
	empl.OrderId = sql.NullInt64{req.Order_id, true}
	err = l.svcCtx.BEmployeeModel.Update(l.ctx, empl)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.AcceptOperationResponse{
		Operation_id: newId,
	}, nil
}
