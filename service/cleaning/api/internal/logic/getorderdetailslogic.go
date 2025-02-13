package logic

import (
	"context"
	"encoding/json"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/category"
	"cleaningservice/service/cleaning/model/contractor"
	"cleaningservice/service/cleaning/model/order"
	"cleaningservice/service/cleaning/model/region"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type GetOrderDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailsLogic {
	return &GetOrderDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderDetailsLogic) GetOrderDetails(req *types.GetOrderDetailsRequest) (resp *types.DetailOrderResponse, err error) {
	_, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, status.Error(401, "Invalid, Not contractor.")
	}

	// Get order details
	order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Get customer details
	customer_item, err := l.svcCtx.BCustomerModel.FindOne(l.ctx, order_item.CustomerId)
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
	address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, order_item.AddressId)
	if err != nil {
		if err == order.ErrNotFound {
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

	// Get contractor details
	// Default contractor details (not found/ haven't been assigned)
	contractor_response := types.DetailContractorResponse{
		Contractor_id:    -1,
		Contractor_photo: "No Contractor Assigned",
		Contractor_name:  "No Contractor Assigned",
		Contractor_type:  "No Contractor Assigned",
		Contact_details:  "No Contractor Assigned",
	}

	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, order_item.ContractorId.Int64)
	if err == nil {
		// Contractor type
		var contractorType string
		if contractor_item.ContractorType == contractor.Employee {
			contractorType = "Employee"
		} else if contractor_item.ContractorType == contractor.Individual {
			contractorType = "Individual"
		}

		contractor_response.Contractor_id = contractor_item.ContractorId
		contractor_response.Contractor_photo = contractor_item.ContractorPhoto.String
		contractor_response.Contractor_name = contractor_item.ContractorName
		contractor_response.Contractor_type = contractorType
		contractor_response.Contact_details = contractor_item.ContactDetails
	} else if err != contractor.ErrNotFound {
		return nil, status.Error(500, err.Error())
	}

	// Get Category Details
	category_item, err := l.svcCtx.BCategoryModel.FindOne(l.ctx, order_item.CategoryId)
	if err != nil {
		if err == category.ErrNotFound {
			return nil, status.Error(404, "Invalid, Category not found.")
		}
		return nil, status.Error(500, err.Error())
	}
	category_response := types.DetailCategoryResponse{
		Category_id:          category_item.CategoryId,
		Category_addr:        category_item.CategoryAddr,
		Category_name:        category_item.CategoryName,
		Category_description: category_item.CategoryDescription,
	}

	// Get Basic Service Details
	var basic_items types.SelectedServiceList
	err = json.Unmarshal([]byte(order_item.BasicItems), &basic_items)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Get Additional Service Details
	var additional_items types.SelectedServiceList
	err = json.Unmarshal([]byte(order_item.AdditionalItems.String), &additional_items)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Create order response
	order_response := types.DetailOrderResponse{
		Order_id:              order_item.OrderId,
		Customer_info:         customer_response,
		Address_info:          address_response,
		Contractor_info:       contractor_response,
		Finance_id:            order_item.FinanceId.Int64,
		Category:              category_response,
		Basic_items:           basic_items,
		Additional_items:      additional_items,
		Order_description:     order_item.OrderDescription.String,
		Order_comments:        order_item.OrderComments.String,
		Current_deposite_rate: int(order_item.CurrentDepositeRate),
		Deposite_amount:       order_item.DepositeAmount,
		Final_amount:          order_item.FinalAmount,
		Item_amount:           order_item.ItemAmount,
		Gst_amount:            order_item.GstAmount,
		Surcharge_item:        order_item.SurchargeItem,
		Surcharge_rate:        int(order_item.SurchargeRate),
		Surcharge_amount:      order_item.SurchargeAmount,
		Surcharge_description: order_item.SurchargeDescription.String,
		Total_amount:          order_item.TotalAmount,
		Balance_amount:        order_item.BalanceAmount,
		Post_date:             order_item.PostDate.Format("2006-01-02 15:04:05"),
		Reserve_date:          order_item.ReserveDate.Format("2006-01-02 15:04:05"),
		Finish_date:           order_item.FinishDate.Time.Format("2006-01-02 15:04:05"),
		Payment_date:          order_item.PaymentDate.Time.Format("2006-01-02 15:04:05"),
		Status:                int(order_item.Status),
		Urgent_flag:           int(order_item.UrgantFlag),
	}

	return &order_response, nil
}
