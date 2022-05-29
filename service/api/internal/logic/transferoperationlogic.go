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
	"cleaningservice/service/model/operation"
	"cleaningservice/service/model/order"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type TransferOperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransferOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferOperationLogic {
	return &TransferOperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransferOperationLogic) TransferOperation(req *types.TransferOperationRequest) (resp *types.TransferOperationResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, status.Error(401, "Invalid, Not contractor.")
	}

	// Valid order
	order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		return nil, status.Error(404, "Invalid, Order not found.")
	}
	if order_item.Status != order.Pending || order_item.Status != order.Working {
		return nil, errorx.NewCodeError(401, "Order is currently unable to be transfer.")
	}

	// Valid contractor
	if uid != order_item.ContractorId.Int64 {
		return nil, status.Error(404, "Invalid, Order not found.")
	}

	// Create operation record
	operation_item := operation.BOperation{
		ContractorId: uid,
		OrderId:      req.Order_id,
		Operation:    operation.Transfer,
		IssueDate:    time.Now(),
	}

	_, err = l.svcCtx.BOperationModel.Insert(l.ctx, &operation_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Transfer
	order_item.ContractorId = sql.NullInt64{0, false}
	order_item.FinanceId = sql.NullInt64{0, false}
	order_item.UrgantFlag = 1
	order_item.Status = order.Transfering
	err = l.svcCtx.BOrderModel.Update(l.ctx, order_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// TODO: Signal tranfer order

	return &types.TransferOperationResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
