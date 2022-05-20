package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/cryptx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/contractor"
	"cleaningservice/service/model/contractorservice"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UpdateContractorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateContractorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContractorLogic {
	return &UpdateContractorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateContractorLogic) UpdateContractor(req *types.UpdateContractorRequest) (resp *types.UpdateContractorResponse, err error) {
	logx.Info("function entrance")
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	var contractorId int64
	if role == variables.Contractor {
		contractorId = uid
	} else if role == variables.Company {
		contractorId = req.Contractor_id
	} else {
		return nil, status.Error(401, "Not Company/Contractor.")
	}

	// Get contractor details
	cont, err := l.svcCtx.BContractorModel.FindOne(l.ctx, contractorId)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, status.Error(404, "Invalid, Contractor not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Verify company and contractor
	if role == variables.Contractor {
		if req.Work_status == variables.Vacant || req.Work_status == variables.InRest {
			cont.WorkStatus = int64(req.Work_status)
		}

		if req.Link_code != "" {
			cont.LinkCode = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Link_code)
		}
	} else if role == variables.Company {
		if cont.FinanceId != uid {
			return nil, status.Error(404, "Invalid, Contractor not found.")
		}
	}

	// Update address details
	if req.Address_info.Street != "No Address" {
		if !cont.AddressId.Valid {
			// Create new address
			newItem := address.BAddress{
				Street:    req.Address_info.Street,
				Suburb:    req.Address_info.Suburb,
				Postcode:  req.Address_info.Postcode,
				City:      req.Address_info.City,
				StateCode: req.Address_info.State_code,
				Country:   req.Address_info.Country,
				Lat:       req.Address_info.Lat,
				Lng:       req.Address_info.Lng,
				Formatted: req.Address_info.Formatted,
			}

			res, err := l.svcCtx.BAddressModel.Insert(l.ctx, &newItem)
			if err != nil {
				return nil, status.Error(500, err.Error())
			}

			newId, err := res.LastInsertId()
			if err != nil {
				return nil, status.Error(500, err.Error())
			}

			cont.AddressId = sql.NullInt64{newId, true}
		} else {
			// Update address
			if err == nil {
				err = l.svcCtx.BAddressModel.Update(l.ctx, &address.BAddress{
					AddressId: cont.AddressId.Int64,
					Street:    req.Address_info.Street,
					Suburb:    req.Address_info.Suburb,
					Postcode:  req.Address_info.Postcode,
					City:      req.Address_info.City,
					StateCode: req.Address_info.State_code,
					Country:   req.Address_info.Country,
					Lat:       req.Address_info.Lat,
					Lng:       req.Address_info.Lng,
					Formatted: req.Address_info.Formatted,
				})

				if err != nil {
					return nil, status.Error(500, err.Error())
				}
			}
		}
	}

	// Update contractor details
	err = l.svcCtx.BContractorModel.Update(l.ctx, &contractor.BContractor{
		ContractorId:    contractorId,
		ContractorPhoto: sql.NullString{req.Contractor_photo, req.Contractor_photo != ""},
		ContractorName:  req.Contractor_name,
		ContractorType:  cont.ContractorType,
		ContactDetails:  req.Contact_details,
		FinanceId:       cont.FinanceId,
		AddressId:       cont.AddressId,
		LinkCode:        cont.LinkCode,
		WorkStatus:      cont.WorkStatus,
		OrderId:         cont.OrderId,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Add new contractor services
	for _, new_service := range req.New_services {
		_, err = l.svcCtx.RContractorServiceModel.Insert(l.ctx, &contractorservice.RContractorService{
			ContractorId: contractorId,
			ServiceId:    new_service,
		})
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	}

	// Remove old contractor services
	for _, old_service := range req.Remove_services {
		err = l.svcCtx.RContractorServiceModel.Delete(l.ctx, contractorId, old_service)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	}

	return &types.UpdateContractorResponse{}, nil
}
