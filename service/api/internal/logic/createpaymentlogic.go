package logic

import (
	"context"
	"database/sql"
	"time"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/payment"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreatePaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePaymentLogic {
	return &CreatePaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePaymentLogic) CreatePayment(req *types.CreatePaymentRequest) (resp *types.CreatePaymentResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)

	exp_time, err := time.Parse("2006-01-02 15:04:05", req.Expiry_time)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	newItem := payment.BPayment{
		CardNumber:   req.Card_number,
		HolderName:   req.Holder_name,
		ExpiryTime:   exp_time,
		SecurityCode: req.Security_code,
	}

	res, err := l.svcCtx.BPaymentModel.Insert(l.ctx, &newItem)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	if role == variables.Company {
		company, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		company.PaymentId = sql.NullInt64{newId, true}

		err = l.svcCtx.BCompanyModel.Update(l.ctx, company)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	}

	return &types.CreatePaymentResponse{
		Payment_id: newId,
	}, nil
}
