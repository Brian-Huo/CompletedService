package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/contractor"
	"cleaningservice/service/model/order"
	"cleaningservice/service/model/orderrecommend"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RecommendOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendOrderLogic {
	return &RecommendOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendOrderLogic) RecommendOrder(req *types.RecommendOrderRequest) (resp *types.RecommendOrderResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, status.Error(401, "Invalid, Not contractor.")
	}

	var orderList []types.DetailOrderResponse
	orderrecommend_items, err := l.svcCtx.ROrderRecommendModel.List(uid)
	if err == nil {
		for _, order_id := range *orderrecommend_items {
			// Get order details
			order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, order_id)
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
				if err == address.ErrNotFound {
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

			order_response := types.DetailOrderResponse{
				Order_id:              order_item.OrderId,
				Customer_info:         customer_response,
				Contractor_info:       contractor_response,
				Address_info:          address_response,
				Finance_id:            order_item.FinanceId.Int64,
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
			}

			orderList = append(orderList, order_response)
		}
	} else if err != orderrecommend.ErrNotFound {
		return nil, status.Error(500, err.Error())
	}

	return &types.RecommendOrderResponse{
		Items: orderList,
	}, nil
}
