package logic

import (
	"context"
	"database/sql"
	"log"
	"time"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/company"
	"cleaningservice/service/model/payment"

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

	newPayment := payment.BPayment{
		CardNumber:   req.Payment_info.Card_number,
		HolderName:   req.Payment_info.Holder_name,
		ExpiryTime:   exp_time,
		SecurityCode: req.Payment_info.Security_code,
	}

	payRes, err := l.svcCtx.BPaymentModel.Insert(l.ctx, &newPayment)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	paymentId, err := payRes.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Create address for company
	newAddress := address.BAddress{
		Street:    req.Address_info.Street,
		Suburb:    req.Address_info.Suburb,
		Postcode:  req.Address_info.Postcode,
		StateCode: req.Address_info.State_code,
		Country:   "AU",
	}

	addressRes, err := l.svcCtx.BAddressModel.Insert(l.ctx, &newAddress)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	addressId, err := addressRes.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Create company details
	newCompany := company.BCompany{
		CompanyName:       req.Company_name,
		PaymentId:         sql.NullInt64{paymentId, true},
		DirectorName:      sql.NullString{req.Director_name, req.Director_name != ""},
		ContactDetails:    req.Contact_details,
		RegisteredAddress: sql.NullInt64{addressId, true},
		DepositeRate:      int64(req.Deposite_rate),
		FinanceStatus:     int64(variables.Active),
	}

	res, err := l.svcCtx.BCompanyModel.Insert(l.ctx, &newCompany)
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
