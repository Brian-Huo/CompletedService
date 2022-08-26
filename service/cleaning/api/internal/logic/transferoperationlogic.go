package logic

import (
	"context"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/order"
	"cleaningservice/service/cleaning/model/orderqueue/transferqueue"

	"github.com/zeromicro/go-zero/core/logx"
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
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, errorx.NewCodeError(401, "Invalid, Not contractor.")
	}

	// Valid order
	order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		return nil, errorx.NewCodeError(404, "Invalid, Order not found.")
	}
	if order_item.Status != order.Pending && order_item.Status != order.Working {
		return nil, errorx.NewCodeError(401, "Order is currently unable to be transfered.")
	}

	// Valid contractor
	if uid != order_item.ContractorId.Int64 {
		return nil, errorx.NewCodeError(404, "Invalid, Order not found.")
	}

	// Get Customer details
	customer_item, err := l.svcCtx.BCustomerModel.FindOne(l.ctx, order_item.CustomerId)
	if err != nil {
		return nil, errorx.NewCodeError(404, "Invalid, Customer not found.")
	}

	// Create operation record
	_, err = l.svcCtx.BOperationModel.RecordTransfer(l.ctx, uid, req.Order_id)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Transfer
	err = l.svcCtx.BOrderModel.Transfer(l.ctx, order_item.OrderId)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	//Record tranfer order
	go l.svcCtx.RTransferQueueModel.Insert(&transferqueue.RTransferQueue{OrderId: req.Order_id, Contact: customer_item.CustomerPhone})

	return &types.TransferOperationResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
