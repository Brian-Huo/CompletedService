package logic

import (
	"context"
	"time"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/orderqueue"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/contractor"
	"cleaningservice/service/cleaning/model/order"
	"cleaningservice/service/cleaning/model/orderqueue/paymentqueue"

	"github.com/zeromicro/go-zero/core/logx"
)

type FinishOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFinishOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FinishOrderLogic {
	return &FinishOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FinishOrderLogic) FinishOrder(req *types.FinishOrderRequest) (resp *types.FinishOrderResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, errorx.NewCodeError(401, "Invalid, Not Contractor.")
	}

	// Get order details
	order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Order not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Validate contractor
	if uid != order_item.ContractorId.Int64 {
		return nil, errorx.NewCodeError(404, "Invalid, Order not found.")
	}

	// Finish order
	if order_item.Status == order.Unpaid || order_item.Status == order.Completed {
		return nil, errorx.NewCodeError(401, "Order has been finished.")
	} else if order_item.Status == order.Cancelled {
		return nil, errorx.NewCodeError(401, "Order has been canceled.")
	} else if order_item.Status == order.Queuing {
		return nil, errorx.NewCodeError(401, "Order is in queue.")
	} else if order_item.Status == order.Pending {
		return nil, errorx.NewCodeError(401, "Order haven't start yet.")
	}

	// Update order details
	err = l.svcCtx.BOrderModel.Finish(l.ctx, req.Order_id)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Get contractor status
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, order_item.ContractorId.Int64)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Validate contractor status
	if contractor_item.WorkStatus == contractor.InWork {
		contractor_item.WorkStatus = contractor.Vacant
	} else {
		return nil, errorx.NewCodeError(401, "Contractor is not in work.")
	}

	// Update contractor details
	err = l.svcCtx.BContractorModel.Update(l.ctx, contractor_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	go orderqueue.PushOne(order_item.OrderId)
	go l.svcCtx.RPaymentQueueModel.Insert(&paymentqueue.RPaymentQueue{
		OrderId: order_item.OrderId,
		DueDate: time.Now().Add(time.Hour * 24 * 8).Format("2006-01-02"),
	})

	return &types.FinishOrderResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
