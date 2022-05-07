package logic

import (
	"context"
	"sort"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/company"
	"cleaningservice/service/model/customerpayment"
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
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)

	if role == variables.Company {
		comp, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
		if err != nil {
			if err == company.ErrNotFound {
				return nil, status.Error(404, "Invalid, Company not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		if comp.PaymentId.Int64 != req.Payment_id {
			return nil, status.Error(404, "Invalid, Payment not found.")
		}
	} else if role == variables.Customer {
		customerPayment, err := l.svcCtx.RCustomerPaymentModel.FindAllByCustomer(l.ctx, uid)
		if err != nil {
			if err == customerpayment.ErrNotFound {
				return nil, status.Error(404, "Invalid, Payment not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		i := sort.Search(len(customerPayment), func(i int) bool { return req.Payment_id <= customerPayment[i].PaymentId })
		if i >= len(customerPayment) || customerPayment[i].PaymentId != req.Payment_id {
			return nil, status.Error(404, "Invalid, Payment not found.")
		}
	}

	res, err := l.svcCtx.BPaymentModel.FindOne(l.ctx, req.Payment_id)
	if err != nil {
		if err == payment.ErrNotFound {
			return nil, status.Error(404, "Invalid, Payment not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &types.DetailPaymentResponse{
		Payment_id:    res.PaymentId,
		Card_number:   res.CardNumber,
		Holder_name:   res.HolderName,
		Expiry_time:   res.ExpiryTime.Format("02/01/2006"),
		Security_code: res.SecurityCode,
	}, nil
}
