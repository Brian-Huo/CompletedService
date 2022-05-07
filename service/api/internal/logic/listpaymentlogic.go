package logic

import (
	"context"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/customerpayment"
	"cleaningservice/service/model/payment"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListPaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPaymentLogic {
	return &ListPaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPaymentLogic) ListPayment(req *types.ListPaymentRequest) (resp *types.ListPaymentResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role != variables.Customer {
		return nil, status.Error(401, "Invalid, Not customer.")
	}

	res, err := l.svcCtx.RCustomerPaymentModel.FindAllByCustomer(l.ctx, uid)
	if err != nil {
		if err == customerpayment.ErrNotFound {
			return nil, status.Error(404, "Invalid, Customer payment not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailPaymentResponse{}

	for _, item := range res {
		pay, err := l.svcCtx.BPaymentModel.FindOne(l.ctx, item.PaymentId)
		if err != nil {
			if err == payment.ErrNotFound {
				return nil, status.Error(404, "Invalid, Payment not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		newItem := types.DetailPaymentResponse{
			Payment_id:    pay.PaymentId,
			Card_number:   pay.CardNumber,
			Holder_name:   pay.HolderName,
			Expiry_time:   pay.ExpiryTime.Format("2006-01-02 15:04:05"),
			Security_code: pay.SecurityCode,
		}

		allItems = append(allItems, newItem)
	}

	return &types.ListPaymentResponse{
		Items: allItems,
	}, nil
}
