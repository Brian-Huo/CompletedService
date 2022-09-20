package logic

import (
	"context"
	"strings"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/address"
	"cleaningservice/service/cleaning/model/region"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAddressLogic {
	return &CreateAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAddressLogic) CreateAddress(req *types.CreateAddressRequest) (resp *types.CreateAddressResponse, err error) {
	// Check address region
	region_item := region.BRegion{
		RegionName: req.Suburb,
		RegionType: "Suburb",
		Postcode:   req.Postcode,
		StateCode:  req.State_code,
		StateName:  req.State_name,
	}
	_, err = l.svcCtx.BRgionModel.Enquire(l.ctx, &region_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	address_item := address.BAddress{
		Street:    req.Street,
		Suburb:    req.Suburb,
		Postcode:  req.Postcode,
		Property:  strings.ToLower(req.Property),
		City:      req.City,
		Lat:       req.Lat,
		Lng:       req.Lng,
		Formatted: req.Formatted,
	}

	res, err := l.svcCtx.BAddressModel.Insert(l.ctx, &address_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateAddressResponse{
		Address_id: newId,
	}, nil
}
