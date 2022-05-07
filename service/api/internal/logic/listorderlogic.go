package logic

import (
	"context"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/order"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrderLogic {
	return &ListOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListOrderLogic) ListOrder(req *types.ListOrderRequest) (resp *types.ListOrderResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role == variables.Employee {
		return nil, status.Error(401, "Invalid, Not customer/company.")
	}

	var res []*order.BOrder
	if role == variables.Customer {
		res, err = l.svcCtx.BOrderModel.FindAllByCustomer(l.ctx, uid)
	} else if role == variables.Company {
		res, err = l.svcCtx.BOrderModel.FindAllByCompany(l.ctx, uid)
	}
	if err != nil {
		if err == order.ErrNotFound {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailOrderResponse{}

	for _, item := range res {
		newItem := types.DetailOrderResponse{
			Order_id:              item.OrderId,
			Customer_id:           item.CustomerId,
			Company_id:            item.CompanyId,
			Address_id:            item.AddressId,
			Design_id:             item.DesignId,
			Deposite_payment:      item.DepositePayment,
			Deposite_amount:       item.DepositeAmount,
			Current_deposite_rate: int(item.CurrentDepositeRate),
			Deposite_date:         item.DepositeDate.Format("2006-01-02 15:04:05"),
			Final_payment:         item.FinalPayment.Int64,
			Final_amount:          item.FinalAmount,
			Final_payment_date:    item.FinalPaymentDate.Time.Format("2006-01-02 15:04:05"),
			Total_fee:             item.TotalFee,
			Order_description:     item.OrderDescription.String,
			Post_date:             item.PostDate.Format("2006-01-02 15:04:05"),
			Reserve_date:          item.ReserveDate.Format("2006-01-02 15:04:05"),
			Finish_date:           item.FinishDate.Time.Format("2006-01-02 15:04:05"),
			Status:                int(item.Status),
		}

		allItems = append(allItems, newItem)
	}

	return &types.ListOrderResponse{
		Items: allItems,
	}, nil
}
