package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/contractor"

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

	// Get contractor details
	var contractor_item *contractor.BContractor
	if role == variables.Company {
		contractor_item, err = l.svcCtx.BContractorModel.FindOne(l.ctx, req.Contractor_id)
		if err != nil {
			if err == contractor.ErrNotFound {
				return nil, status.Error(404, "Invalid, Contractor not found.")
			}
			return nil, status.Error(500, err.Error())
		}
	} else if role == variables.Contractor {
		contractor_item, err = l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
		if err != nil {
			if err == contractor.ErrNotFound {
				return nil, status.Error(404, "Invalid, Contractor not found.")
			}
			return nil, status.Error(500, err.Error())
		}
	} else {
		return nil, status.Error(401, "Invalid, Unauthorized action.")
	}

	// Contractor type
	var contractorType string
	if contractor_item.ContractorType == contractor.Employee {
		contractorType = "Employee"
	} else if contractor_item.ContractorType == contractor.Individual {
		contractorType = "Individual"
	}

	// Contractor address details
	address_response := types.DetailAddressResponse{
		Address_id: -1,
		Street:     "No Address",
	}

	// Get address details
	address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, contractor_item.AddressId.Int64)
	if err == nil {
		region_item, err := l.svcCtx.BRgionModel.FindOneByPostcode(l.ctx, address_item.Postcode)
		if err == nil {
			address_response.Address_id = address_item.AddressId
			address_response.Street = address_item.Street
			address_response.Suburb = address_item.Suburb
			address_response.Postcode = address_item.Postcode
			address_response.Property = address_item.Property
			address_response.City = address_item.City
			address_response.State_code = region_item.StateCode
			address_response.State_name = region_item.StateName
			address_response.Lat = address_item.Lat
			address_response.Lng = address_item.Lng
			address_response.Formatted = address_item.Formatted
		}
	}

	// Get category details
	category_list, err := l.svcCtx.RSubscriptionModel.ListSubscribeGroup(l.ctx, uid)
	if err != nil {
		return nil, status.Error(404, "Invalid, Category list not found.")
	}

	return &types.DetailContractorResponse{
		Contractor_id:    contractor_item.ContractorId,
		Contractor_photo: contractor_item.ContractorPhoto.String,
		Contractor_name:  contractor_item.ContractorName,
		Contractor_type:  contractorType,
		Contact_details:  contractor_item.ContactDetails,
		Address_info:     address_response,
		Finance_id:       contractor_item.FinanceId,
		Link_code:        contractor_item.LinkCode,
		Work_status:      int(contractor_item.WorkStatus),
		Order_id:         0,
		Category_list:    *category_list,
	}, nil
}
