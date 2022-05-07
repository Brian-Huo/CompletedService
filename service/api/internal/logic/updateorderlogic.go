package logic

import (
	"context"
	"database/sql"
	"time"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UpdateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderLogic) UpdateOrder(req *types.UpdateOrderRequest) (resp *types.UpdateOrderResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Customer {
		return nil, status.Errorf(401, "Invalid, Not customer.")
	}

	// Get origin order
	res, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(404, "Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}
	if res.CustomerId != uid {
		return nil, status.Error(401, "Invalid customer order id.")
	}

	reserve_date, err := time.Parse("2006-01-02 15:04:05", req.Reserve_date)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// if modify reserve_date
	if reserve_date != res.ReserveDate {
		// if reserve_date is close to current time in 12 hours, return error
		if time.Now().Add(time.Hour * 12).Before(reserve_date) {
			return nil, status.Error(500, "Reserve date is futher less than 12 hours.")
		}
		res.ReserveDate = reserve_date
	} else {
		// if pay the final ammout
		_, err := l.svcCtx.BPaymentModel.FindOne(l.ctx, req.Final_payment)
		if err != nil {
			return nil, status.Error(404, "Final_payment not found.")
		}

		res.FinalPayment = sql.NullInt64{req.Final_payment, true}
		res.FinalPaymentDate = sql.NullTime{time.Now(), true}
		res.OrderDescription = sql.NullString{req.Order_description, true}
		res.FinishDate = sql.NullTime{time.Now(), true}
	}

	err = l.svcCtx.BOrderModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &types.UpdateOrderResponse{}, nil
}
