package logic

import (
	"context"
	"time"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/operation"

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

	_, err = l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, status.Error(404, "Invalid, Contractor not found.")
	}

	l.receiveOrder(uid, req.Order_id)

	ord, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		return nil, status.Error(404, "Invalid, Order not found.")
	}

	if ord.Status != int64(variables.Queuing) {
		return nil, errorx.NewCodeError(401, "Order is currently unavailable.")
	}

	newItem := operation.BOperation{
		ContractorId: uid,
		OrderId:      req.Order_id,
		Operation:    int64(variables.Decline),
		IssueDate:    time.Now(),
	}

	res, err := l.svcCtx.BOperationModel.Insert(l.ctx, &newItem)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.DeclineOperationResponse{
		Operation_id: newId,
	}, nil
}

func (l *DeclineOperationLogic) receiveOrder(contractorId int64, orderId int64) {
	go l.svcCtx.BScheduleModel.Delete(contractorId, orderId)
}
