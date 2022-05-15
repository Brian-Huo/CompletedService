package logic

import (
	"context"
	"time"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/operation"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateOperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOperationLogic {
	return &CreateOperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOperationLogic) CreateOperation(req *types.CreateOperationRequest) (resp *types.CreateOperationResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role != variables.Employee {
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
		Operation:  req.Operation,
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

	return &types.CreateOperationResponse{
		Operation_id: newId,
	}, nil

}
