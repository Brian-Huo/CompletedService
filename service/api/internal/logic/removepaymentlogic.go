package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/company"
	"cleaningservice/service/model/customerpayment"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemovePaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemovePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemovePaymentLogic {
	return &RemovePaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemovePaymentLogic) RemovePayment(req *types.RemovePaymentRequest) (resp *types.RemovePaymentResponse, err error) {
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

		comp.PaymentId = sql.NullInt64{0, false}

		err = l.svcCtx.BCompanyModel.Update(l.ctx, comp)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	} else if role == variables.Customer {
		err := l.svcCtx.RCustomerPaymentModel.Delete(l.ctx, uid, req.Payment_id)
		if err != nil {
			if err == customerpayment.ErrNotFound {
				return nil, status.Error(404, "Invalid, Customer payment record not found.")
			}
			return nil, status.Error(500, err.Error())
		}
	}

	go l.svcCtx.BPaymentModel.Delete(l.ctx, req.Payment_id)

	return
}
