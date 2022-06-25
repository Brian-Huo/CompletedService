package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/cryptx"
	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/address"
	"cleaningservice/service/cleaning/model/contractor"

	"github.com/zeromicro/go-zero/core/logx"
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
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	}

	var contractorId int64
	if role == variables.Contractor {
		contractorId = uid
	} else if role == variables.Company {
		contractorId = req.Contractor_id
	} else {
		return nil, errorx.NewCodeError(401, "Not Company/Contractor.")
	}

	// Get contractor details
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, contractorId)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Contractor not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Verify company and contractor
	if role == variables.Contractor {
		if req.Work_status == int(contractor.Vacant) || req.Work_status == int(contractor.InRest) {
			contractor_item.WorkStatus = int64(req.Work_status)
		}

		if req.Link_code != "" {
			contractor_item.LinkCode = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Link_code)
		}
	} else if role == variables.Company {
		if contractor_item.FinanceId != uid {
			return nil, errorx.NewCodeError(404, "Invalid, Contractor not found.")
		}
	}

	// Update address details
	if req.Address_info.Street != "No Address" && req.Address_info.Address_id != -1 {
		if !contractor_item.AddressId.Valid {
			// Create new address
			address_struct := address.BAddress{
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

			res, err := l.svcCtx.BAddressModel.Insert(l.ctx, &address_struct)
			if err != nil {
				return nil, errorx.NewCodeError(500, err.Error())
			}

			newId, err := res.LastInsertId()
			if err != nil {
				return nil, errorx.NewCodeError(500, err.Error())
			}

			contractor_item.AddressId = sql.NullInt64{newId, true}
		} else {
			// Update address
			if err == nil {
				err = l.svcCtx.BAddressModel.Update(l.ctx, &address.BAddress{
					AddressId: contractor_item.AddressId.Int64,
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
					return nil, errorx.NewCodeError(500, err.Error())
				}
			}
		}
	}

	// Update contractor details
	err = l.svcCtx.BContractorModel.Update(l.ctx, &contractor.BContractor{
		ContractorId:    contractorId,
		ContractorPhoto: sql.NullString{req.Contractor_photo, req.Contractor_photo != ""},
		ContractorName:  req.Contractor_name,
		ContractorType:  contractor_item.ContractorType,
		ContactDetails:  req.Contact_details,
		FinanceId:       contractor_item.FinanceId,
		AddressId:       contractor_item.AddressId,
		LinkCode:        contractor_item.LinkCode,
		WorkStatus:      contractor_item.WorkStatus,
		OrderId:         contractor_item.OrderId,
		CategoryList:    contractor_item.CategoryList,
	})
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	return &types.UpdateContractorResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
