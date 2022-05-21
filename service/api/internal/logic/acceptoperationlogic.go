package logic

import (
	"context"
	"database/sql"
	"time"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/contractor"
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
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, status.Error(401, "Invalid, Not contractor.")
	}

	// Check contractor status
	cont, err := l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, status.Error(404, "Invalid, Contractor not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if cont.WorkStatus != int64(variables.Vacant) {
		return nil, errorx.NewCodeError(401, "Invalid, Contractor should not double accept order(s).")
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
		return nil, errorx.NewCodeError(401, "Order is currently unavailable.")
	}

	// Create operaction records
	newItem := operation.BOperation{
		ContractorId: uid,
		OrderId:      req.Order_id,
		Operation:    int64(variables.Accept),
		IssueDate:    time.Now(),
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
	ord.ContractorId = sql.NullInt64{uid, true}
	ord.FinanceId = sql.NullInt64{cont.FinanceId, true}
	ord.Status = int64(variables.Working)
	err = l.svcCtx.BOrderModel.Update(l.ctx, ord)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Update contractor details
	cont.WorkStatus = int64(variables.InWork)
	cont.OrderId = sql.NullInt64{req.Order_id, true}
	err = l.svcCtx.BContractorModel.Update(l.ctx, cont)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	l.receiveOrder(uid, ord.OrderId)

	return &types.AcceptOperationResponse{
		Operation_id: newId,
	}, nil
}

func (l *AcceptOperationLogic) receiveOrder(contractorId int64, orderId int64) {
	go l.svcCtx.BScheduleModel.Delete(contractorId, orderId)
}
