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
	res, err := l.svcCtx.BOrderModel.FindOne(l.ctx, req.Order_id)
	if err != nil {
		if err == order.ErrNotFound {
			return nil, status.Error(404, "Invalid, Order not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Get customer details
	cus, err := l.svcCtx.BCustomerModel.FindOne(l.ctx, res.CustomerId)
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
	addr, err := l.svcCtx.BAddressModel.FindOne(l.ctx, res.AddressId)
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
		City:       addr.City,
		State_code: addr.StateCode,
		Country:    addr.Country,
		Lat:        addr.Lat,
		Lng:        addr.Lng,
		Formatted:  addr.Formatted,
	}

	// Get contractor details
	// Default contractor details (not found/ haven't been assigned)
	newCont := types.DetailContractorResponse{
		Contractor_id:    -1,
		Contractor_photo: "No Contractor Assigned",
		Contractor_name:  "No Contractor Assigned",
		Contractor_type:  "No Contractor Assigned",
		Contact_details:  "No Contractor Assigned",
	}

	cont, err := l.svcCtx.BContractorModel.FindOne(l.ctx, res.ContractorId.Int64)
	if err == nil {
		// Contractor type
		var contractorType string
		if cont.ContractorType == int64(variables.Employee) {
			contractorType = "Employee"
		} else if cont.ContractorType == int64(variables.Individual) {
			contractorType = "Individual"
		}

		newCont.Contractor_id = cont.ContractorId
		newCont.Contractor_photo = cont.ContractorPhoto.String
		newCont.Contractor_name = cont.ContractorName
		newCont.Contractor_type = contractorType
		newCont.Contact_details = cont.ContactDetails
		newCont.Finance_id = -1
		newCont.Work_status = -1
		newCont.Order_id = -1
	} else if err != contractor.ErrNotFound {
		return nil, status.Error(500, err.Error())
	}

	order_item := types.GetOrderDetailsResponse{
		Order_id:              res.OrderId,
		Customer_info:         newCus,
		Contractor_info:       newCont,
		Address_info:          newAddr,
		Finance_id:            res.FinanceId.Int64,
		Service_list:          res.ServiceList,
		Deposite_payment:      res.DepositePayment,
		Deposite_amount:       res.DepositeAmount,
		Current_deposite_rate: int(res.CurrentDepositeRate),
		Deposite_date:         res.DepositeDate.Format("2006-01-02 15:04:05"),
		Final_payment:         res.FinalPayment.Int64,
		Final_amount:          res.FinalAmount,
		Final_payment_date:    res.FinalPaymentDate.Time.Format("2006-01-02 15:04:05"),
		Gst_amount:            res.GstAmount,
		Total_fee:             res.TotalFee,
		Order_description:     res.OrderDescription.String,
		Post_date:             res.PostDate.Format("2006-01-02 15:04:05"),
		Reserve_date:          res.ReserveDate.Format("2006-01-02 15:04:05"),
		Finish_date:           res.FinishDate.Time.Format("2006-01-02 15:04:05"),
		Status:                int(res.Status),
	}

	// Replace blank info
	if !res.FinalPayment.Valid {
		order_item.Final_payment = -1
		order_item.Final_payment_date = ""
	}
	if !res.FinishDate.Valid {
		order_item.Finish_date = ""
	}

	return &order_item, nil
}
