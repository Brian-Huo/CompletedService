package logic

import (
	"context"
	"encoding/json"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/address"
	"cleaningservice/service/cleaning/model/category"
	"cleaningservice/service/cleaning/model/contractor"
	"cleaningservice/service/cleaning/model/order"
	"cleaningservice/service/cleaning/model/region"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type GetContractorHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContractorHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContractorHistoryLogic {
	return &GetContractorHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContractorHistoryLogic) GetContractorHistory(req *types.GetContractorHistoryRequest) (resp *types.GetContractorHistoryResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, status.Error(401, "Invalid, Not contractor.")
	}

	orderList := []types.DetailOrderResponse{}
	items, err := l.svcCtx.BOrderModel.ListContractorHistories(l.ctx, uid)
	if err != nil {
		if err == order.ErrNotFound {
			return &types.GetContractorHistoryResponse{
				Items: orderList,
			}, nil
		}
		return nil, status.Error(500, err.Error())
	}

	// Get contractor details
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
	if err != nil {
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

	for _, item := range items {
		// Get customer details
		customer_item, err := l.svcCtx.BCustomerModel.FindOne(l.ctx, item.CustomerId)
		if err != nil {
			if err == order.ErrNotFound {
				return nil, status.Error(404, "Invalid, Customer not found.")
			}
			return nil, status.Error(500, err.Error())
		}
		customer_response := types.DetailCustomerResponse{
			Customer_id:    customer_item.CustomerId,
			Customer_name:  customer_item.CustomerName,
			Customer_phone: customer_item.CustomerPhone,
			Customer_email: customer_item.CustomerEmail,
			Country_code:   customer_item.CountryCode,
		}

		// Get address details
		address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, item.AddressId)
		if err != nil {
			if err == address.ErrNotFound {
				return nil, status.Error(404, "Invalid, Address not found.")
			}
			return nil, status.Error(500, err.Error())
		}
		// Get region details
		region_item, err := l.svcCtx.BRgionModel.FindOneByPostcode(l.ctx, address_item.Postcode)
		if err != nil {
			if err == region.ErrNotFound {
				return nil, status.Error(404, "Invalid, Region not found.")
			}
			return nil, status.Error(500, err.Error())
		}
		address_response := types.DetailAddressResponse{
			Address_id: address_item.AddressId,
			Street:     address_item.Street,
			Suburb:     address_item.Suburb,
			Postcode:   address_item.Postcode,
			Property:   address_item.Property,
			City:       address_item.City,
			State_code: region_item.StateCode,
			State_name: region_item.StateName,
			Lat:        address_item.Lat,
			Lng:        address_item.Lng,
			Formatted:  address_item.Formatted,
		}

		// Get Category Details
		category_item, err := l.svcCtx.BCategoryModel.FindOne(l.ctx, item.CategoryId)
		if err != nil {
			if err == category.ErrNotFound {
				return nil, status.Error(404, "Invalid, Category not found.")
			}
			return nil, status.Error(500, err.Error())
		}
		category_response := types.DetailCategoryResponse{
			Category_id:          category_item.CategoryId,
			Category_name:        category_item.CategoryName,
			Category_description: category_item.CategoryDescription,
		}

		// Get Basic Service Details
		var basic_items types.SelectedServiceStructure
		err = json.Unmarshal([]byte(item.BasicItems), &basic_items)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		// Get Additional Service Details
		var additional_items types.SelectedServiceList
		err = json.Unmarshal([]byte(item.AdditionalItems.String), &additional_items)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		// Create order response
		order_response := types.DetailOrderResponse{
			Order_id:              item.OrderId,
			Customer_info:         customer_response,
			Address_info:          address_response,
			Contractor_info:       contractor_response,
			Finance_id:            item.FinanceId.Int64,
			Category:              category_response,
			Basic_items:           basic_items,
			Additional_items:      additional_items,
			Order_description:     item.OrderDescription.String,
			Order_comments:        item.OrderComments.String,
			Current_deposite_rate: int(item.CurrentDepositeRate),
			Deposite_amount:       item.DepositeAmount,
			Final_amount:          item.FinalAmount,
			Item_amount:           item.ItemAmount,
			Gst_amount:            item.GstAmount,
			Surcharge_item:        item.SurchargeItem,
			Surcharge_rate:        int(item.SurchargeRate),
			Surcharge_amount:      item.SurchargeAmount,
			Total_amount:          item.TotalAmount,
			Balance_amount:        item.BalanceAmount,
			Post_date:             item.PostDate.Format("2006-01-02 15:04:05"),
			Reserve_date:          item.ReserveDate.Format("2006-01-02 15:04:05"),
			Finish_date:           item.FinishDate.Time.Format("2006-01-02 15:04:05"),
			Payment_date:          item.PaymentDate.Time.Format("2006-01-02 15:04:05"),
			Status:                int(item.Status),
			Urgent_flag:           int(item.UrgantFlag),
		}

		orderList = append(orderList, order_response)
	}

	return &types.GetContractorHistoryResponse{
		Items: orderList,
	}, nil
}
