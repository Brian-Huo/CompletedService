package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/employee"
	"cleaningservice/service/model/order"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DetailOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailOrderLogic {
	return &DetailOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailOrderLogic) DetailOrder(req *types.DetailOrderRequest) (resp *types.DetailOrderResponse, err error) {
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
		Address_id:      addr.AddressId,
		Address_details: addr.AddressDetails,
		Suburb:          addr.Suburb,
		Postcode:        addr.Postcode,
		State_code:      addr.StateCode,
		Country:         addr.Country.String,
	}

	// Get employee details
	// Default employee details (not found/ haven't been assigned)
	newEmpl := types.DetailEmployeeResponse{
		Employee_id:     0,
		Employee_photo:  "Employee Not Found",
		Employee_name:   "Employee Not Found",
		Contact_details: "Employee Not Found",
	}

	empl, err := l.svcCtx.BEmployeeModel.FindOne(l.ctx, res.EmployeeId.Int64)
	if err == nil {
		newEmpl.Employee_id = empl.EmployeeId
		newEmpl.Employee_photo = empl.EmployeePhoto.String
		newEmpl.Employee_name = empl.EmployeeName
		newEmpl.Contact_details = empl.ContactDetails
	} else if err != employee.ErrNotFound {
		return nil, status.Error(500, err.Error())
	}

	return &types.DetailOrderResponse{
		Order_id:              res.OrderId,
		Customer_info:         newCus,
		Employee_info:         newEmpl,
		Address_info:          newAddr,
		Company_id:            res.CompanyId.Int64,
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
	}, nil
}
