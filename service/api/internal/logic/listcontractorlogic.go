package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/contractor"

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

	res, err := l.svcCtx.BContractorModel.FindAllByFinance(l.ctx, uid)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, status.Error(404, "Invalid, Contractor not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailContractorResponse{}

	for _, item := range res {
		// Check if contractor resigned
		if item.WorkStatus == int64(variables.Resigned) {
			continue
		}

		// Get all contractor service
		service_list := types.ListContractorServiceResponse{}
		service_res, err := l.svcCtx.RContractorServiceModel.FindAllByContractor(l.ctx, item.ContractorId)
		if err == nil {
			allServices := []types.DetailServiceResponse{}
			for _, res_item := range service_res {
				service_item, err := l.svcCtx.BServiceModel.FindOne(l.ctx, res_item.ServiceId)
				if err != nil {
					break
				}

				service := types.DetailServiceResponse{
					Service_id:          service_item.ServiceId,
					Service_type:        service_item.ServiceType,
					Service_description: service_item.ServiceDescription,
					Service_price:       service_item.ServicePrice,
				}

				allServices = append(allServices, service)
			}
			service_list.Items = allServices
		}

		// Contractor type
		var contractorType string
		if item.ContractorType == int64(variables.Employee) {
			contractorType = "Employee"
		} else if item.ContractorType == int64(variables.Individual) {
			contractorType = "Individual"
		}

		// Contractor address details
		newAddr := types.DetailAddressResponse{
			Address_id: -1,
			Street:     "Not Found",
			Suburb:     "Not Found",
			Postcode:   "Not Found",
			City:       "Not Found",
			State_code: "Not Found",
			Country:    "Not Found",
			Lng:        -1,
			Lat:        -1,
			Formatted:  "Not Found",
		}

		address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, item.AddressId.Int64)
		if err == nil {
			newAddr.Address_id = address_item.AddressId
			newAddr.Street = address_item.Street
			newAddr.Suburb = address_item.Suburb
			newAddr.Postcode = address_item.Postcode
			newAddr.City = address_item.City
			newAddr.State_code = address_item.StateCode
			newAddr.Country = address_item.Country
			newAddr.Lat = address_item.Lat
			newAddr.Lng = address_item.Lng
			newAddr.Formatted = address_item.Formatted
		}

		// Get contractor details
		newItem := types.DetailContractorResponse{
			Contractor_id:      item.ContractorId,
			Contractor_photo:   item.ContractorPhoto.String,
			Contractor_name:    item.ContractorName,
			Contractor_type:    contractorType,
			Contact_details:    item.ContactDetails,
			Address_info:       newAddr,
			Finance_id:         item.FinanceId,
			Link_code:          item.LinkCode,
			Work_status:        int(item.WorkStatus),
			Order_id:           item.OrderId.Int64,
			Contractor_service: service_list,
		}

		allItems = append(allItems, newItem)
	}

	return &types.ListContractorResponse{
		Items: allItems,
	}, nil
}
