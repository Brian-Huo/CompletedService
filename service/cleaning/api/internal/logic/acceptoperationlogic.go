package logic

import (
	"context"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/orderqueue"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/contractor"
	"cleaningservice/service/cleaning/model/order"

	"github.com/zeromicro/go-zero/core/logx"
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
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, errorx.NewCodeError(401, "Invalid, Not contractor.")
	}

	// Check contractor status
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Contractor not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Check order status
	order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Order not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Validate accept operation
	if contractor_item.WorkStatus == contractor.Await {
		return nil, errorx.NewCodeError(401, "Invalid, Contractor haven't registered.")
	}
	if order_item.Status != order.Queuing && order_item.Status != order.Transfering {
		return nil, errorx.NewCodeError(401, "Order is currently unavailable.")
	}

	// Create operaction records
	_, err = l.svcCtx.BOperationModel.RecordAccept(l.ctx, uid, req.Order_id)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Update order details
	err = l.svcCtx.BOrderModel.Accept(l.ctx, req.Order_id, uid, contractor_item.FinanceId)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	l.removeBroadcast(order_item.CategoryId, order_item.OrderId)

	return &types.AcceptOperationResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}

func (l *AcceptOperationLogic) removeBroadcast(groupId int64, orderId int64) {
	go l.svcCtx.BBroadcastModel.Delete(groupId, orderId)
	go orderqueue.Delete(orderId)
}
