package logic

import (
	"context"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"
	"cleaningservice/service/model/order"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DetailOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailOrderLogic {
	return &DetailOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailOrderLogic) DetailOrder(req *types.DetailOrderRequest) (resp *types.DetailOrderResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)

	res, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if role == variables.Customer {
		if uid != res.CustomerId {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
	} else if role == variables.Company {
		if uid != res.CompanyId {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
	} else if role == variables.Employee {
		emp, err := l.svcCtx.BEmployeeModel.FindOne(l.ctx, uid)
		if err != nil {
			if err == employee.ErrNotFound {
				return nil, status.Error(404, "Invalid, Employee not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		if emp.CompanyId != res.CompanyId {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
	}

	return &types.DetailOrderResponse{
		Order_id:              res.OrderId,
		Customer_id:           res.CustomerId,
		Company_id:            res.CompanyId,
		Address_id:            res.AddressId,
		Design_id:             res.DesignId,
		Deposite_payment:      res.DepositePayment,
		Deposite_amount:       res.DepositeAmount,
		Current_deposite_rate: int(res.CurrentDepositeRate),
		Deposite_date:         string(res.DepositeDate.Format("2006-01-02 15:04:05")),
		Final_payment:         res.FinalPayment.Int64,
		Final_amount:          res.FinalAmount,
		Final_payment_date:    res.FinalPaymentDate.Time.Format("2006-01-02 15:04:05"),
		Total_fee:             res.TotalFee,
		Order_description:     res.OrderDescription.String,
		Post_date:             res.PostDate.Format("2006-01-02 15:04:05"),
		Reserve_date:          res.ReserveDate.Format("2006-01-02 15:04:05"),
		Finish_date:           res.FinishDate.Time.Format("2006-01-02 15:04:05"),
		Status:                int(res.Status),
	}, nil
}
