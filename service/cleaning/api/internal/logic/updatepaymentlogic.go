package logic

import (
	"context"
	"time"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/company"
	"cleaningservice/service/cleaning/model/payment"

	"github.com/zeromicro/go-zero/core/logx"
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
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	} else if role != variables.Company {
		return nil, errorx.NewCodeError(401, "Invalid, Not company.")
	}

	company_item, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == company.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Company not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	if company_item.PaymentId.Int64 != req.Payment_id {
		return nil, errorx.NewCodeError(404, "Company payment record not found.")
	}

	expiryTime, err := time.Parse("2006-01-02 15:04:05", req.Expiry_time)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
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
			return nil, errorx.NewCodeError(404, "Payment not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	return &types.UpdatePaymentResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
