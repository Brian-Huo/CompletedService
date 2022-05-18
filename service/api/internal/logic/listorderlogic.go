package logic

import (
	"context"

	"cleaningservice/common/jwtx"
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
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role == variables.Employee {
		return nil, status.Error(401, "Invalid, Not customer/company.")
	}

	// Get all order list
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
		// Get customer details
		cus, err := l.svcCtx.BCustomerModel.FindOne(l.ctx, item.CustomerId)
		if err != nil {
			if err == order.ErrNotFound {
				return nil, status.Error(404, "Invalid, Customer not found.")
			}
			return nil, status.Error(500, err.Error())
		}
		newCus := types.DetailCustomerResponse{
			Customer_id:     cus.CustomerId,
			Customer_name:   cus.CustomerName,
			Contact_details: cus.ContactDetails,
			Country_code:    cus.CountryCode,
		}

		// Get address details
		addr, err := l.svcCtx.BAddressModel.FindOne(l.ctx, item.AddressId)
		if err != nil {
			if err == order.ErrNotFound {
				return nil, status.Error(404, "Invalid, Address not found.")
			}
			return nil, status.Error(500, err.Error())
		}
		newAddr := types.DetailAddressResponse{
			Address_id: addr.AddressId,
			Street:     addr.Street,
			Suburb:     addr.Suburb,
			Postcode:   addr.Postcode,
			State_code: addr.StateCode,
			Country:    addr.Country,
		}

		// Get employee details
		empl, err := l.svcCtx.BEmployeeModel.FindOne(l.ctx, item.EmployeeId.Int64)
		if err != nil {
			if err == order.ErrNotFound {
				return nil, status.Error(404, "Invalid, Employee not found.")
			}
			return nil, status.Error(500, err.Error())
		}
		newEmpl := types.DetailEmployeeResponse{
			Employee_id:     empl.EmployeeId,
			Employee_photo:  empl.EmployeePhoto.String,
			Employee_name:   empl.EmployeeName,
			Contact_details: empl.ContactDetails,
		}

		// Create order detail response
		newItem := types.DetailOrderResponse{
			Order_id:              item.OrderId,
			Customer_info:         newCus,
			Employee_info:         newEmpl,
			Address_info:          newAddr,
			Company_id:            item.CompanyId.Int64,
			Service_list:          item.ServiceList,
			Deposite_payment:      item.DepositePayment,
			Deposite_amount:       item.DepositeAmount,
			Current_deposite_rate: int(item.CurrentDepositeRate),
			Deposite_date:         item.DepositeDate.Format("2006-01-02 15:04:05"),
			Final_payment:         item.FinalPayment.Int64,
			Final_amount:          item.FinalAmount,
			Final_payment_date:    item.FinalPaymentDate.Time.Format("2006-01-02 15:04:05"),
			Gst_amount:            item.GstAmount,
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
