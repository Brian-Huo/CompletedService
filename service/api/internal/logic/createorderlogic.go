package logic

import (
	"context"
	"database/sql"
	"time"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/order"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderRequest) (resp *types.CreateOrderResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role == variables.Customer {
		return nil, status.Error(401, "Invalid, Not customer.")
	}

	// Exist detail check
	_, err = l.svcCtx.BCustomerModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, status.Error(404, "Invalid, Customer not found.")
	}
	_, err = l.svcCtx.BCompanyModel.FindOne(l.ctx, req.Company_id)
	if err != nil {
		return nil, status.Error(404, "Invalid, Company not found.")
	}
	_, err = l.svcCtx.BAddressModel.FindOne(l.ctx, req.Address_id)
	if err != nil {
		return nil, status.Error(404, "Invalid, Address not found.")
	}
	_, err = l.svcCtx.BDesignModel.FindOne(l.ctx, req.Design_id)
	if err != nil {
		return nil, status.Error(404, "Invalid, Design not found.")
	}
	_, err = l.svcCtx.BPaymentModel.FindOne(l.ctx, req.Deposite_payment)
	if err != nil {
		return nil, status.Error(404, "Invalid, Payment detail not found.")
	}

	// Get time variables
	deposite_date, err := time.Parse("2006-01-02 15:04:05", req.Deposite_date)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	final_payment_date, err := time.Parse("2006-01-02 15:04:05", req.Final_payment_date)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	reserve_date, err := time.Parse("2006-01-02 15:04:05", req.Reserve_date)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newItem := order.BOrder{
		CustomerId:       uid,
		CompanyId:        req.Company_id,
		AddressId:        req.Address_id,
		DesignId:         req.Design_id,
		DepositePayment:  req.Deposite_payment,
		DepositeDate:     deposite_date,
		FinalPayment:     sql.NullInt64{req.Final_payment, true},
		FinalPaymentDate: sql.NullTime{final_payment_date, true},
		OrderDescription: sql.NullString{req.Order_description, true},
		PostDate:         time.Now(),
		ReserveDate:      reserve_date,
	}

	res, err := l.svcCtx.BOrderModel.Insert(l.ctx, &newItem)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateOrderResponse{
		Order_id: newId,
	}, nil
}
