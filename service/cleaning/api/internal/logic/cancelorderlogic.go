package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/errorx"
	"cleaningservice/common/orderqueue"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/contractor"
	"cleaningservice/service/cleaning/model/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOrderLogic {
	return &CancelOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelOrderLogic) CancelOrder(req *types.CancelOrderRequest) (resp *types.CancelOrderResponse, err error) {
	// Update order details
	order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Order not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	err = l.svcCtx.BOrderModel.Cancel(l.ctx, order_item.OrderId)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	l.removeBroadcast(order_item.CategoryId, order_item.OrderId)

	// Contractor workstatus update
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, order_item.ContractorId.Int64)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Contractor not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	contractor_item.WorkStatus = contractor.Vacant
	contractor_item.OrderId = sql.NullInt64{Int64: 0, Valid: false}

	err = l.svcCtx.BContractorModel.Update(l.ctx, contractor_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	return &types.CancelOrderResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}

func (l *CancelOrderLogic) removeBroadcast(groupId int64, orderId int64) {
	go l.svcCtx.BBroadcastModel.Delete(groupId, orderId)
	go orderqueue.Delete(orderId)
}
