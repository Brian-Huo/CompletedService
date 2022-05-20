package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/company"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UpdateAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddressLogic {
	return &UpdateAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAddressLogic) UpdateAddress(req *types.UpdateAddressRequest) (resp *types.UpdateAddressResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	// check address id vaild for company
	if role != variables.Company {
		return nil, status.Error(401, "Invalid, Not company.")
	}

	res, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == company.ErrNotFound {
			return nil, status.Error(404, "Invalid, Company not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if res.RegisteredAddress.Valid && res.RegisteredAddress.Int64 != req.Address_id {
		return nil, status.Error(401, "Invalid company address id.")
	}

	err = l.svcCtx.BAddressModel.Update(l.ctx, &address.BAddress{
		AddressId: req.Address_id,
		Street:    req.Street,
		Suburb:    req.Suburb,
		Postcode:  req.Postcode,
		City:      req.City,
		StateCode: req.State_code,
		Country:   req.Country,
		Lat:       req.Lat,
		Lng:       req.Lng,
		Formatted: req.Formatted,
	})
	if err != nil {
		if err == address.ErrNotFound {
			return nil, status.Error(404, "Invalid, Address not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &types.UpdateAddressResponse{}, nil
}
