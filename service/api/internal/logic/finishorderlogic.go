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

	ord, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if uid != ord.ContractorId.Int64 {
		return nil, status.Error(404, "Invalid, Order not found.")
	}
	logx.Info("finishing order")
	// Finish order
	if ord.Status == order.Working {
		ord.Status = order.Unpaid
	} else {
		return nil, errorx.NewCodeError(401, "Order cannot be finished twice.")
	}

	err = l.svcCtx.BOrderModel.Update(l.ctx, ord)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Update contractor status
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, ord.ContractorId.Int64)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	contractor_item.WorkStatus = contractor.Vacant
	contractor_item.OrderId = sql.NullInt64{0, false}

	err = l.svcCtx.BContractorModel.Update(l.ctx, contractor_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.FinishOrderResponse{}, nil
}
