package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/contractor"
	"cleaningservice/service/model/order"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, status.Error(401, "Invalid, Not Contractor.")
	}

	// Get order details
	order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Validate contractor
	if uid != order_item.ContractorId.Int64 {
		return nil, status.Error(404, "Invalid, Order not found.")
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
	order_item.Status = order.Unpaid

	// Get contractor status
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, order_item.ContractorId.Int64)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Validate contractor status
	if contractor_item.WorkStatus == contractor.InWork {
		contractor_item.WorkStatus = contractor.Vacant
		contractor_item.OrderId = sql.NullInt64{0, false}
	} else {
		return nil, errorx.NewCodeError(401, "Contractor is not in work.")
	}

	// Update order details
	err = l.svcCtx.BOrderModel.Update(l.ctx, order_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Update contractor details
	err = l.svcCtx.BContractorModel.Update(l.ctx, contractor_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.FinishOrderResponse{}, nil
}
