package logic

import (
	"context"
	"database/sql"
	"time"

	"cleaningservice/common/errorx"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/order"
	"cleaningservice/service/cleaning/model/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOrderLogic {
	return &PayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayOrderLogic) PayOrder(req *types.PayOrderRequest) (resp *types.PayOrderResponse, err error) {
	// Check order status
	order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Order not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}
	if order_item.Status == order.Completed {
		return nil, errorx.NewCodeError(401, "Invalid, Order has been paid.")
	}

	// Pay order here

	// Pay order here

	// Create or find payment details
	exp_time, err := time.Parse("2006-01-02 15:04:05", req.Final_info.Expiry_time)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	var paymentId int64
	res, err := l.svcCtx.BPaymentModel.FindOneByCardNumber(l.ctx, req.Final_info.Card_number)
	if err == payment.ErrNotFound {
		payment_item, err := l.svcCtx.BPaymentModel.Insert(l.ctx, &payment.BPayment{
			CardNumber:   req.Final_info.Card_number,
			HolderName:   req.Final_info.Holder_name,
			ExpiryTime:   exp_time,
			SecurityCode: req.Final_info.Security_code,
		})
		if err != nil {
			return nil, errorx.NewCodeError(500, err.Error())
		}

		paymentId, err = payment_item.LastInsertId()
		if err != nil {
			return nil, errorx.NewCodeError(500, err.Error())
		}
	} else if err == nil {
		paymentId = res.PaymentId
	} else {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Update order details
	order_item.FinalPayment = sql.NullInt64{paymentId, true}
	order_item.Status = order.Completed

	err = l.svcCtx.BOrderModel.Update(l.ctx, order_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	return &types.PayOrderResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
