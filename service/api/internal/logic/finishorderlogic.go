package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
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
	} else if role != variables.Employee {
		return nil, status.Error(401, "Invalid, Not employee.")
	}

	ord, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if uid != ord.EmployeeId.Int64 {
		return nil, status.Error(404, "Invalid, Order not found.")
	}

	// Finish order
	ord.Status = int64(variables.Unpaid)

	err = l.svcCtx.BOrderModel.Update(l.ctx, ord)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Update employee status
	empl, err := l.svcCtx.BEmployeeModel.FindOne(l.ctx, ord.EmployeeId.Int64)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	empl.WorkStatus = int64(variables.Vacant)
	empl.OrderId = sql.NullInt64{0, false}

	err = l.svcCtx.BEmployeeModel.Update(l.ctx, empl)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return
}
