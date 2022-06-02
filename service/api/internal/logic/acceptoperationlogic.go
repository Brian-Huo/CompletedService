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
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, status.Error(404, "Invalid, Contractor not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Check order status
	order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Validate accept operation
	if contractor_item.WorkStatus == contractor.Await {
		return nil, errorx.NewCodeError(401, "Invalid, Contractor haven't registed.")
	}
	if order_item.Status != order.Queuing {
		return nil, errorx.NewCodeError(401, "Order is currently unavailable.")
	}

	// Create operaction records
	newItem := operation.BOperation{
		ContractorId: uid,
		OrderId:      req.Order_id,
		Operation:    operation.Accept,
		IssueDate:    time.Now(),
	}

	_, err = l.svcCtx.BOperationModel.Insert(l.ctx, &newItem)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Update order details
	order_item.ContractorId = sql.NullInt64{uid, true}
	order_item.FinanceId = sql.NullInt64{contractor_item.FinanceId, true}
	order_item.Status = order.Pending
	err = l.svcCtx.BOrderModel.Update(l.ctx, order_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	l.removeBroadcast(order_item.CategoryId, order_item.OrderId)

	return &types.AcceptOperationResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}

func (l *AcceptOperationLogic) removeBroadcast(groupId int64, orderId int64) {
	go l.svcCtx.BBroadcastModel.Delete(groupId, orderId)
}
