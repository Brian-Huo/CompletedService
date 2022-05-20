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

	// Get all contractor service
	service_list := types.ListContractorServiceResponse{}
	service_res, err := l.svcCtx.RContractorServiceModel.FindAllByContractor(l.ctx, res.ContractorId)
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
	if res.ContractorType == int64(variables.Employee) {
		contractorType = "Employee"
	} else if res.ContractorType == int64(variables.Individual) {
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

	address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, res.AddressId.Int64)
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

	return &types.DetailContractorResponse{
		Contractor_id:      res.ContractorId,
		Contractor_photo:   res.ContractorPhoto.String,
		Contractor_name:    res.ContractorName,
		Contractor_type:    contractorType,
		Contact_details:    res.ContactDetails,
		Address_info:       newAddr,
		Finance_id:         res.FinanceId,
		Link_code:          res.LinkCode,
		Work_status:        int(res.WorkStatus),
		Order_id:           res.OrderId.Int64,
		Contractor_service: service_list,
	}, nil
}
