package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/contractor"
	"cleaningservice/service/cleaning/model/region"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListContractorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListContractorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListContractorLogic {
	return &ListContractorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListContractorLogic) ListContractor(req *types.ListContractorRequest) (resp *types.ListContractorResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, Json format error")
	} else if role != variables.Company {
		return nil, status.Error(401, "Invalid, Not company.")
	}

	contractor_items, err := l.svcCtx.BContractorModel.FindAllByFinance(l.ctx, uid)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, status.Error(404, "Invalid, Contractor not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailContractorResponse{}

	for _, item := range contractor_items {
		// Check if contractor resigned
		if item.WorkStatus == contractor.Resigned {
			continue
		}

		// Contractor type
		var contractorType string
		if item.ContractorType == contractor.Employee {
			contractorType = "Employee"
		} else if item.ContractorType == contractor.Individual {
			contractorType = "Individual"
		}

		// Contractor address details
		address_response := types.DetailAddressResponse{
			Address_id: -1,
			Street:     "Not Found",
			Suburb:     "Not Found",
			Postcode:   "Not Found",
			City:       "Not Found",
			State_code: "Not Found",
			Lng:        -1,
			Lat:        -1,
			Formatted:  "Not Found",
		}

		address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, item.AddressId.Int64)
		if err == nil {
			// Get region details
			region_item, err := l.svcCtx.BRgionModel.FindOneByPostcode(l.ctx, address_item.Postcode)
			if err != nil {
				if err == region.ErrNotFound {
					return nil, status.Error(404, "Invalid, Region not found.")
				}
				return nil, status.Error(500, err.Error())
			}

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

		// Get category details
		category_list, err := l.svcCtx.RSubscriptionModel.ListSubscribeGroup(l.ctx, item.ContractorId)
		if err != nil {
			return nil, status.Error(404, "Invalid, Category list not found.")
		}

		// Get contractor current working order
		cur_order, _ := l.svcCtx.BOrderModel.FindCurrentWorkingOneByContractor(l.ctx, item.ContractorId)

		// Get contractor details
		contractor_response := types.DetailContractorResponse{
			Contractor_id:    item.ContractorId,
			Contractor_photo: item.ContractorPhoto.String,
			Contractor_name:  item.ContractorName,
			Contractor_type:  contractorType,
			Contact_details:  item.ContactDetails,
			Address_info:     address_response,
			Finance_id:       item.FinanceId,
			Link_code:        item.LinkCode,
			Work_status:      int(item.WorkStatus),
			Order_id:         cur_order,
			Category_list:    *category_list,
		}

		allItems = append(allItems, contractor_response)
	}

	return &types.ListContractorResponse{
		Items: allItems,
	}, nil
}
