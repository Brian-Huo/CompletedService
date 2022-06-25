package logic

import (
	"context"
	"time"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/operation"
	"cleaningservice/service/cleaning/model/order"
	"cleaningservice/service/cleaning/model/orderdelay"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DeclineOperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeclineOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeclineOperationLogic {
	return &DeclineOperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeclineOperationLogic) DeclineOperation(req *types.DeclineOperationRequest) (resp *types.DeclineOperationResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, status.Error(401, "Invalid, Not contractor.")
	}

	// Validate contractor details
	_, err = l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, status.Error(404, "Invalid, Contractor not found.")
	}

	l.receiveOrder(uid, req.Order_id)

	// Validate order details
	order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		return nil, status.Error(404, "Invalid, Order not found.")
	}

	if order_item.Status != order.Queuing {
		return nil, errorx.NewCodeError(401, "Order is currently unavailable.")
	}

	newItem := operation.BOperation{
		ContractorId: uid,
		OrderId:      req.Order_id,
		Operation:    operation.Decline,
		IssueDate:    time.Now(),
	}

	_, err = l.svcCtx.BOperationModel.Insert(l.ctx, &newItem)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.DeclineOperationResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}

func (l *DeclineOperationLogic) receiveOrder(contractorId int64, order_itemerId int64) {
	go l.svcCtx.ROrderDelayModel.Insert(&orderdelay.ROrderDelay{
		ContractorId: contractorId,
		OrderId:      order_itemerId,
	})
}
