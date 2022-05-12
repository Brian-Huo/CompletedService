package logic

import (
	"context"
	"time"

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
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Employee {
		return nil, status.Error(401, "Invalid, Not employee.")
	}

	_, err = l.svcCtx.BEmployeeModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, status.Error(404, "Invalid, Employee not found.")
	}

	_, err = l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		return nil, status.Error(404, "Invalid, Order not found.")
	}

	newItem := operation.BOperation{
		EmployeeId: uid,
		OrderId:    req.Order_id,
		Operation:  int64(variables.Decline),
		IssueDate:  time.Now(),
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
	return
}
