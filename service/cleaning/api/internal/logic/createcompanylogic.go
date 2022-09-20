package logic

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

	"cleaningservice/common/jwtx"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/address"
	"cleaningservice/service/cleaning/model/company"
	"cleaningservice/service/cleaning/model/payment"
	"cleaningservice/service/cleaning/model/region"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCompanyLogic {
	return &CreateCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCompanyLogic) CreateCompany(req *types.CreateCompanyRequest) (resp *types.CreateCompanyResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role != 100 && uid != 0 {
		log.Println("Backend broken, security leak...")
		return nil, status.Error(500, err.Error())
	}

	// Create Payment details for company
	exp_time, err := time.Parse("2006-01-02 15:04:05", req.Payment_info.Expiry_time)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	payment_item := payment.BPayment{
		CardNumber:   req.Payment_info.Card_number,
		HolderName:   req.Payment_info.Holder_name,
		ExpiryTime:   exp_time,
		SecurityCode: req.Payment_info.Security_code,
	}

	payRes, err := l.svcCtx.BPaymentModel.Insert(l.ctx, &payment_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	paymentId, err := payRes.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Create address for company
	// Check address region
	region_item := region.BRegion{
		RegionName: req.Address_info.Suburb,
		RegionType: "Suburb",
		Postcode:   req.Address_info.Postcode,
		StateCode:  req.Address_info.State_code,
		StateName:  req.Address_info.State_name,
	}
	_, err = l.svcCtx.BRgionModel.Enquire(l.ctx, &region_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	address_item := address.BAddress{
		Street:    req.Address_info.Street,
		Suburb:    req.Address_info.Suburb,
		Postcode:  req.Address_info.Postcode,
		Property:  strings.ToLower(req.Address_info.Property),
		City:      req.Address_info.City,
		Lat:       req.Address_info.Lat,
		Lng:       req.Address_info.Lng,
		Formatted: req.Address_info.Formatted,
	}

	addressRes, err := l.svcCtx.BAddressModel.Insert(l.ctx, &address_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	addressId, err := addressRes.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Create company details
	company_item := company.BCompany{
		CompanyName:       req.Company_name,
		PaymentId:         sql.NullInt64{Int64: paymentId, Valid: true},
		DirectorName:      sql.NullString{String: req.Director_name, Valid: req.Director_name != ""},
		ContactDetails:    req.Contact_details,
		RegisteredAddress: addressId,
		DepositeRate:      int64(req.Deposite_rate),
		FinanceStatus:     company.Active,
	}

	res, err := l.svcCtx.BCompanyModel.Insert(l.ctx, &company_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateCompanyResponse{
		Company_id: newId,
	}, nil
}
