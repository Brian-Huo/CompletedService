package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/address"
	"cleaningservice/service/cleaning/model/region"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DetailAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailAddressLogic {
	return &DetailAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailAddressLogic) DetailAddress(req *types.DetailAddressRequest) (resp *types.DetailAddressResponse, err error) {
	address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, req.Address_id)
	if err != nil {
		if err == address.ErrNotFound {
			return nil, status.Error(404, "Invalid, Address not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	region_item, err := l.svcCtx.BRgionModel.FindOneByPostcode(l.ctx, address_item.Postcode)
	if err != nil {
		if err == region.ErrNotFound {
			return nil, status.Error(404, "Invalid, Region not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &types.DetailAddressResponse{
		Street:     address_item.Street,
		Suburb:     address_item.Suburb,
		Postcode:   address_item.Postcode,
		Property:   address_item.Property,
		City:       address_item.City,
		State_code: region_item.StateCode,
		State_name: region_item.StateName,
		Lat:        address_item.Lat,
		Lng:        address_item.Lng,
		Formatted:  address_item.Formatted,
	}, nil
}
