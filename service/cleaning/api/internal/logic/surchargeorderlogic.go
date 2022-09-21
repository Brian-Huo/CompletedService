package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"cleaningservice/common/errorx"
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

type SurchargeOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSurchargeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SurchargeOrderLogic {
	return &SurchargeOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SurchargeOrderLogic) SurchargeOrder(req *types.SurchargeOrderRequest) (resp *types.SurchargeOrderResponse, err error) {
	logx.Info("function entrance\n")
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, errorx.NewCodeError(401, "Invalid, Not contractor.")
	}

	// Get order details
	order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Order not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Validate contractor
	if uid != order_item.ContractorId.Int64 {
		return nil, errorx.NewCodeError(401, "Invalid, Not assigned contractor.")
	}

	// Surchharge Order
	order_item.SurchargeItem = req.Surcharge_item
	order_item.SurchargeRate = int64(req.Surcharge_rate*variables.Surcharge_factor - 10)
	order_item.SurchargeAmount = order_item.ItemAmount * float64(req.Surcharge_rate) / 100
	order_item.SurchargeDescription = sql.NullString{String: req.Surcharge_description, Valid: len(req.Surcharge_description) != 0}
	order_item.GstAmount = (order_item.ItemAmount + order_item.SurchargeAmount) / variables.GST
	order_item.TotalAmount = order_item.ItemAmount + order_item.SurchargeAmount + order_item.GstAmount

	err = l.svcCtx.BOrderModel.Update(l.ctx, order_item)
	if err != nil {
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
		Total_amount:          order_item.TotalAmount,
		Balance_amount:        order_item.BalanceAmount,
		Post_date:             order_item.PostDate.Format("2006-01-02 15:04:05"),
		Reserve_date:          order_item.ReserveDate.Format("2006-01-02 15:04:05"),
		Finish_date:           order_item.FinishDate.Time.Format("2006-01-02 15:04:05"),
		Payment_date:          order_item.PaymentDate.Time.Format("2006-01-02 15:04:05"),
		Status:                int(order_item.Status),
		Urgent_flag:           int(order_item.UrgantFlag),
	}

	return &types.SurchargeOrderResponse{
		Code: 200,
		Msg:  "success",
		Data: order_response,
	}, nil
}
