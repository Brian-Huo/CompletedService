package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/contractor"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DetailContractorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailContractorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailContractorLogic {
	return &DetailContractorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailContractorLogic) DetailContractor(req *types.DetailContractorRequest) (resp *types.DetailContractorResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	var res *contractor.BContractor
	if role != variables.Contractor {
		res, err = l.svcCtx.BContractorModel.FindOne(l.ctx, req.Contractor_id)
		if err != nil {
			if err == contractor.ErrNotFound {
				return nil, status.Error(404, "Invalid, Contractor not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		if role == variables.Customer {
			res.LinkCode = ""
			res.OrderId = sql.NullInt64{0, false}
			res.WorkStatus = -1
		}
	} else if role == variables.Contractor {
		res, err = l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
		if err != nil {
			if err == contractor.ErrNotFound {
				return nil, status.Error(404, "Invalid, Contractor not found.")
			}
			return nil, status.Error(500, err.Error())
		}
	}

	// Contractor type
	var contractorType string
	if res.ContractorType == contractor.Employee {
		contractorType = "Employee"
	} else if res.ContractorType == contractor.Individual {
		contractorType = "Individual"
	}

	// Contractor address details
	address_response := types.DetailAddressResponse{
		Address_id: -1,
		Street:     "No Address",
	}

	// Get address details
	address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, res.AddressId.Int64)
	if err == nil {
		address_response.Address_id = address_item.AddressId
		address_response.Street = address_item.Street
		address_response.Suburb = address_item.Suburb
		address_response.Postcode = address_item.Postcode
		address_response.City = address_item.City
		address_response.State_code = address_item.StateCode
		address_response.Country = address_item.Country
		address_response.Lat = address_item.Lat
		address_response.Lng = address_item.Lng
		address_response.Formatted = address_item.Formatted
	}

	return &types.DetailContractorResponse{
		Contractor_id:    res.ContractorId,
		Contractor_photo: res.ContractorPhoto.String,
		Contractor_name:  res.ContractorName,
		Contractor_type:  contractorType,
		Contact_details:  res.ContactDetails,
		Address_info:     address_response,
		Finance_id:       res.FinanceId,
		Link_code:        res.LinkCode,
		Work_status:      int(res.WorkStatus),
		Order_id:         res.OrderId.Int64,
	}, nil
}
