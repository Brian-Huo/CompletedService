package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/company"
	"cleaningservice/service/model/payment"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DetailPaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailPaymentLogic {
	return &DetailPaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailPaymentLogic) DetailPayment(req *types.DetailPaymentRequest) (resp *types.DetailPaymentResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	if role == variables.Company {
		company_item, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
		if err != nil {
			if err == company.ErrNotFound {
				return nil, status.Error(404, "Invalid, Company not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		if company_item.PaymentId.Int64 != req.Payment_id {
			return nil, status.Error(404, "Invalid, Payment not found.")
		}
	}

	payment_item, err := l.svcCtx.BPaymentModel.FindOne(l.ctx, req.Payment_id)
	if err != nil {
		if err == payment.ErrNotFound {
			return nil, status.Error(404, "Invalid, Payment not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &types.DetailPaymentResponse{
		Payment_id:    payment_item.PaymentId,
		Card_number:   payment_item.CardNumber,
		Holder_name:   payment_item.HolderName,
		Expiry_time:   payment_item.ExpiryTime.Format("02/01/2006"),
		Security_code: payment_item.SecurityCode,
	}, nil
}
