package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/contractor"
	"cleaningservice/service/model/subscriberecord"

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

	var contractor_item *contractor.BContractor
	if role != variables.Contractor {
		contractor_item, err = l.svcCtx.BContractorModel.FindOne(l.ctx, req.Contractor_id)
		if err != nil {
			if err == contractor.ErrNotFound {
				return nil, status.Error(404, "Invalid, Contractor not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		if role == variables.Customer {
			contractor_item.LinkCode = ""
			contractor_item.OrderId = sql.NullInt64{0, false}
			contractor_item.WorkStatus = -1
		}
	} else if role == variables.Contractor {
		contractor_item, err = l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
		if err != nil {
			if err == contractor.ErrNotFound {
				return nil, status.Error(404, "Invalid, Contractor not found.")
			}
			return nil, status.Error(500, err.Error())
		}
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

	// Get category details
	subscriberecord_items, err := l.svcCtx.RSubscribeRecordModel.FindAllByContractorId(l.ctx, uid)
	if err != nil {
		if err == subscriberecord.ErrNotFound {
			return nil, status.Error(404, "Invalid, Category not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	var category_list []int64
	for _, item := range subscriberecord_items {
		group_item, err := l.svcCtx.BSubscribeGroupModel.FindOne(l.ctx, item.GroupId)
		if err != nil {
			continue
		}

		category_list = append(category_list, group_item.Category)
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
		Order_id:         contractor_item.OrderId.Int64,
		Category_list:    category_list,
	}, nil
}
