package logic

import (
	"context"
	"database/sql"
	"strings"

	"cleaningservice/common/cryptx"
	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/address"
	"cleaningservice/service/cleaning/model/contractor"
	"cleaningservice/service/cleaning/model/region"

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

	// Check region details
	_, err = l.svcCtx.BRgionModel.Enquire(l.ctx, &region.BRegion{
		RegionName: req.Address_info.Suburb,
		RegionType: "Suburb",
		Postcode:   req.Address_info.Postcode,
		StateCode:  req.Address_info.State_code,
		StateName:  req.Address_info.State_name,
	})
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}
	// Update address details
	if !contractor_item.AddressId.Valid {
		// Create new address
		address_struct := address.BAddress{
			Street:    req.Address_info.Street,
			Suburb:    req.Address_info.Suburb,
			Postcode:  req.Address_info.Postcode,
			Property:  strings.ToLower(req.Address_info.Property),
			City:      req.Address_info.City,
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

		contractor_item.AddressId = sql.NullInt64{Int64: newId, Valid: true}
	} else {
		// Update address
		if err == nil {
			err = l.svcCtx.BAddressModel.Update(l.ctx, &address.BAddress{
				AddressId: contractor_item.AddressId.Int64,
				Street:    req.Address_info.Street,
				Suburb:    req.Address_info.Suburb,
				Postcode:  req.Address_info.Postcode,
				Property:  req.Address_info.Property,
				City:      req.Address_info.City,
				Lat:       req.Address_info.Lat,
				Lng:       req.Address_info.Lng,
				Formatted: req.Address_info.Formatted,
			})
			if err != nil {
				return nil, errorx.NewCodeError(500, err.Error())
			}
		}
	}

	// Update contractor details
	err = l.svcCtx.BContractorModel.Update(l.ctx, &contractor.BContractor{
		ContractorId:    contractorId,
		ContractorPhoto: sql.NullString{String: req.Contractor_photo, Valid: req.Contractor_photo != ""},
		ContractorName:  req.Contractor_name,
		ContractorType:  contractor_item.ContractorType,
		ContactDetails:  req.Contact_details,
		FinanceId:       contractor_item.FinanceId,
		AddressId:       contractor_item.AddressId,
		LinkCode:        contractor_item.LinkCode,
		WorkStatus:      contractor_item.WorkStatus,
	})
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	return &types.UpdateContractorResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
