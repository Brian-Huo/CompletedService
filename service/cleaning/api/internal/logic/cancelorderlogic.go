package logic

import (
	"context"

	"cleaningservice/common/errorx"
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

	// Contractor workstatus update
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, order_item.ContractorId.Int64)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Contractor not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	if order_item.Status == order.Working {
		contractor_item.WorkStatus = contractor.Vacant
	}

	err = l.svcCtx.BOrderModel.Cancel(l.ctx, order_item.OrderId)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	l.removeBroadcast(order_item.CategoryId, order_item.OrderId)

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
	go l.svcCtx.RAwaitQueueModel.Delete(orderId)
	go l.svcCtx.RTransferQueueModel.Delete(orderId)
}
