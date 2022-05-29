package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/category"
	"cleaningservice/service/model/contractor"
	"cleaningservice/service/model/order"

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

func (l *GetOrderDetailsLogic) GetOrderDetails(req *types.GetOrderDetailsRequest) (resp *types.GetOrderDetailsResponse, err error) {
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
		Customer_id:     customer_item.CustomerId,
		Customer_name:   customer_item.CustomerName,
		Contact_details: customer_item.ContactDetails,
		Country_code:    customer_item.CountryCode,
	}

	// Get address details
	address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, order_item.AddressId)
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
		contractor_response.Finance_id = -1
		contractor_response.Work_status = -1
		contractor_response.Order_id = -1
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
		Category_name:        category_item.CategoryName,
		Category_description: category_item.CategoryDescription,
	}

	order_response := types.GetOrderDetailsResponse{
		Order_id:              order_item.OrderId,
		Customer_info:         customer_response,
		Contractor_info:       contractor_response,
		Address_info:          address_response,
		Finance_id:            order_item.FinanceId.Int64,
		Category:              category_response,
		Service_list:          order_item.ServiceList,
		Deposite_payment:      order_item.DepositePayment,
		Deposite_amount:       order_item.DepositeAmount,
		Current_deposite_rate: int(order_item.CurrentDepositeRate),
		Deposite_date:         order_item.DepositeDate.Format("2006-01-02 15:04:05"),
		Final_payment:         order_item.FinalPayment.Int64,
		Final_amount:          order_item.FinalAmount,
		Final_payment_date:    order_item.FinalPaymentDate.Time.Format("2006-01-02 15:04:05"),
		Gst_amount:            order_item.GstAmount,
		Total_fee:             order_item.TotalFee,
		Order_description:     order_item.OrderDescription.String,
		Post_date:             order_item.PostDate.Format("2006-01-02 15:04:05"),
		Reserve_date:          order_item.ReserveDate.Format("2006-01-02 15:04:05"),
		Finish_date:           order_item.FinishDate.Time.Format("2006-01-02 15:04:05"),
		Status:                int(order_item.Status),
		Urgent_flag:           int(order_item.UrgantFlag),
	}

	// Replace blank info
	if !order_item.FinalPayment.Valid {
		order_response.Final_payment = -1
		order_response.Final_payment_date = ""
	}
	if !order_item.FinishDate.Valid {
		order_response.Finish_date = ""
	}

	return &order_response, nil
}
