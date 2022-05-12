package logic

import (
	"context"
	"time"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/company"
	"cleaningservice/service/model/payment"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UpdatePaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePaymentLogic {
	return &UpdatePaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePaymentLogic) UpdatePayment(req *types.UpdatePaymentRequest) (resp *types.UpdatePaymentResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)

	if role != variables.Company {
		return nil, status.Error(401, "Invalid, Not company.")
	}

	comp, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == company.ErrNotFound {
			return nil, status.Error(404, "Company not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if comp.PaymentId.Int64 != req.Payment_id {
		return nil, status.Error(404, "Company payment record not found.")
	}

	expiryTime, err := time.Parse("2006-01-02 15:04:05", req.Expiry_time)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.BPaymentModel.Update(l.ctx, &payment.BPayment{
		PaymentId:    req.Payment_id,
		CardNumber:   req.Card_number,
		HolderName:   req.Holder_name,
		ExpiryTime:   expiryTime,
		SecurityCode: req.Security_code,
	})
	if err != nil {
		if err == payment.ErrNotFound {
			return nil, status.Error(404, "Payment not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &types.UpdatePaymentResponse{}, nil
}
