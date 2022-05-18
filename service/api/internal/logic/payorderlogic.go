package logic

import (
	"context"
	"database/sql"
	"time"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/order"
	"cleaningservice/service/model/payment"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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
	ord, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}
	if ord.Status == int64(variables.Completed) {
		return nil, status.Error(401, "Invalid, Order has been paid.")
	}

	// Pay order here

	// Pay order here

	// Create or find payment details
	exp_time, err := time.Parse("2006-01-02 15:04:05", req.Final_info.Expiry_time)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	var paymentId int64
	pay, err := l.svcCtx.BPaymentModel.FindOneByCardNumber(l.ctx, req.Final_info.Card_number)
	if err == payment.ErrNotFound {
		res, err := l.svcCtx.BPaymentModel.Insert(l.ctx, &payment.BPayment{
			CardNumber:   req.Final_info.Card_number,
			HolderName:   req.Final_info.Holder_name,
			ExpiryTime:   exp_time,
			SecurityCode: req.Final_info.Security_code,
		})

		paymentId, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	} else if err == nil {
		paymentId = pay.PaymentId
	} else {
		return nil, status.Error(500, err.Error())
	}

	// Update order details
	ord.FinalPayment = sql.NullInt64{paymentId, true}
	ord.Status = int64(variables.Completed)

	err = l.svcCtx.BOrderModel.Update(l.ctx, ord)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return
}
