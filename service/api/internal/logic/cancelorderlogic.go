package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/broadcast"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"
	"cleaningservice/service/model/order"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CancelOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOrderLogic {
	return &CancelOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelOrderLogic) CancelOrder(req *types.CancelOrderRequest) (resp *types.CancelOrderResponse, err error) {
	ord, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	ord.Status = int64(variables.Cancelled)

	err = l.svcCtx.BOrderModel.Update(l.ctx, ord)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	empl, err := l.svcCtx.BEmployeeModel.FindOne(l.ctx, ord.EmployeeId)
	if err != nil {
		if err == employee.ErrNotFound {
			return nil, status.Error(404, "Invalid, Employee not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	empl.WorkStatus = int64(variables.Vacant)
	empl.OrderId = sql.NullInt64{0, false}

	err = l.svcCtx.BEmployeeModel.Update(l.ctx, empl)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	go broadcast.OneBroadcast()

	return
}
