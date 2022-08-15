package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/address"
	"cleaningservice/service/cleaning/model/contractor"
	"cleaningservice/service/cleaning/model/order"
	"cleaningservice/util"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type StartOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartOrderLogic {
	return &StartOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartOrderLogic) StartOrder(req *types.StartOrderRequest) (resp *types.StartOrderResponse, err error) {
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
	address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, order_item.AddressId)
	if err != nil {
		if err == address.ErrNotFound {
			return nil, status.Error(404, "Invalid, Address not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Get contractor status
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, order_item.ContractorId.Int64)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Validate contractor
	if uid != order_item.ContractorId.Int64 {
		return nil, status.Error(404, "Invalid, Order not found.")
	}

	// Validate contractor status
	if contractor_item.WorkStatus == contractor.Vacant {
		contractor_item.WorkStatus = contractor.InWork
		contractor_item.OrderId = sql.NullInt64{Int64: req.Order_id, Valid: true}
	} else {
		return nil, errorx.NewCodeError(401, "Contractor is not vacant.")
	}

	// validate order status
	if order_item.Status == order.Unpaid || order_item.Status == order.Completed {
		return nil, errorx.NewCodeError(401, "Order has been finished.")
	} else if order_item.Status == order.Cancelled {
		return nil, errorx.NewCodeError(401, "Order has been canceled.")
	} else if order_item.Status == order.Queuing {
		return nil, errorx.NewCodeError(401, "Order is in queue.")
	} else if order_item.Status == order.Working {
		return nil, errorx.NewCodeError(401, "Order is in work.")
	}

	// Validate contractor is available for this order
	if util.CheckPointsDistance(req.Lat, req.Lng, address_item.Lat, address_item.Lng, variables.Inwork_distance) {
		order_item.Status = order.Working
	} else {
		return nil, errorx.NewCodeError(401, "Contractor is not available for this order.")
	}

	// Update order status
	err = l.svcCtx.BOrderModel.Update(l.ctx, order_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Update contractor status
	err = l.svcCtx.BContractorModel.Update(l.ctx, contractor_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.StartOrderResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
