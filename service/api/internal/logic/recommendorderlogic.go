package logic

import (
	"context"
	"strconv"
	"strings"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/category"
	"cleaningservice/service/model/contractor"
	"cleaningservice/service/model/order"

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

	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, status.Error(404, "Contractor not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	var orderList []types.DetailOrderResponse
	category_list := strings.Split(contractor_item.CategoryList.String, variables.Separator)
	for _, group_str := range category_list {
		// Get all broadcasting orders from group
		group_id, err := strconv.ParseInt(group_str, 10, 64)
		if err != nil {
			logx.Error("contractor category wrong format %s", group_str)
			continue
		}
		order_list, err := l.svcCtx.BBroadcastModel.List(group_id)
		if err != nil {
			logx.Error("get broadcasting orders from group %s failed", group_id)
			continue
		}

		// Get order details
		for _, order_id := range *order_list {
			order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, order_id)
			if err != nil {
				go l.svcCtx.BBroadcastModel.Delete(group_id, order_id)
			}

			// Get customer details
			customer_item, err := l.svcCtx.BCustomerModel.FindOne(l.ctx, order_item.CustomerId)
			if err != nil {
				if err == order.ErrNotFound {
					return nil, status.Error(404, "Invalid, Customer not found.")
				}
				return nil, status.Error(500, err.Error())
			}

			// Get address details
			address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, order_item.AddressId)
			if err != nil {
				if err == address.ErrNotFound {
					return nil, status.Error(404, "Invalid, Address not found.")
				}
				return nil, status.Error(500, err.Error())
			}

			// Get category details
			category_item, err := l.svcCtx.BCategoryModel.FindOne(l.ctx, order_item.CategoryId)
			if err != nil {
				if err == category.ErrNotFound {
					return nil, status.Error(404, "Invalid, Category not found.")
				}
				return nil, status.Error(500, err.Error())
			}

			// Construct order response
			order_response := types.DetailOrderResponse{
				Order_id: order_item.OrderId,
				Customer_info: types.DetailCustomerResponse{
					Customer_id:     customer_item.CustomerId,
					Customer_name:   customer_item.CustomerName,
					Contact_details: customer_item.ContactDetails,
					Country_code:    customer_item.CountryCode,
				},
				Address_info: types.DetailAddressResponse{
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
				},
				Finance_id: order_item.FinanceId.Int64,
				Category: types.DetailCategoryResponse{
					Category_id:          category_item.CategoryId,
					Category_name:        category_item.CategoryName,
					Category_description: category_item.CategoryDescription,
				},
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

			orderList = append(orderList, order_response)
		}
	}

	return &types.RecommendOrderResponse{
		Items: orderList,
	}, nil
}
