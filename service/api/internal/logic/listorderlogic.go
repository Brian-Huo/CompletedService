package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/contractor"
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
	}

	// Get all order list
	var order_items []*order.BOrder
	if role == variables.Customer {
		order_items, err = l.svcCtx.BOrderModel.FindAllByCustomer(l.ctx, uid)
	} else if role == variables.Company {
		order_items, err = l.svcCtx.BOrderModel.FindAllByFinance(l.ctx, uid)
	} else {
		return nil, status.Error(401, "Invalid, Unauthorised action.")
	}

	if err != nil {
		if err == order.ErrNotFound {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailOrderResponse{}
	for _, item := range order_items {
		// Get customer details
		customer_item, err := l.svcCtx.BCustomerModel.FindOne(l.ctx, item.CustomerId)
		if err != nil {
			if err == order.ErrNotFound {
				return nil, status.Error(404, "Invalid, Customer not found.")
			}
			return nil, status.Error(500, err.Error())
		}
		customer_response := types.DetailCustomerResponse{
			Customer_id:     customer_item.CustomerId,
			Customer_name:   customer_item.CustomerName,
			Contact_details: customer_item.ContactDetails,
			Country_code:    customer_item.CountryCode,
		}

		// Get address details
		address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, item.AddressId)
		if err != nil {
			if err == order.ErrNotFound {
				return nil, status.Error(404, "Invalid, Address not found.")
			}
			return nil, status.Error(500, err.Error())
		}
		address_response := types.DetailAddressResponse{
			Address_id: address_item.AddressId,
			Street:     address_item.Street,
			Suburb:     address_item.Suburb,
			Postcode:   address_item.Postcode,
			City:       address_item.City,
			State_code: address_item.StateCode,
			Country:    address_item.Country,
			Lat:        address_item.Lat,
			Lng:        address_item.Lng,
			Formatted:  address_item.Formatted,
		}

		// Get contractor details
		contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, item.ContractorId.Int64)
		if err != nil {
			if err == order.ErrNotFound {
				return nil, status.Error(404, "Invalid, Contractor not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		// Contractor type
		var contractorType string
		if contractor_item.ContractorType == contractor.Employee {
			contractorType = "Employee"
		} else if contractor_item.ContractorType == contractor.Individual {
			contractorType = "Individual"
		}

		contractor_response := types.DetailContractorResponse{
			Contractor_id:    contractor_item.ContractorId,
			Contractor_photo: contractor_item.ContractorPhoto.String,
			Contractor_name:  contractor_item.ContractorName,
			Contractor_type:  contractorType,
			Contact_details:  contractor_item.ContactDetails,
		}

		// Create order detail response
		order_response := types.DetailOrderResponse{
			Order_id:              item.OrderId,
			Customer_info:         customer_response,
			Contractor_info:       contractor_response,
			Address_info:          address_response,
			Finance_id:            item.FinanceId.Int64,
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

		allItems = append(allItems, order_response)
	}

	return &types.ListOrderResponse{
		Items: allItems,
	}, nil
}
