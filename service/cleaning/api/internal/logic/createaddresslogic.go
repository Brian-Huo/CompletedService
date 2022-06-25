package logic

import (
	"context"

	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/address"

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
	newItem := address.BAddress{
		Street:    req.Street,
		Suburb:    req.Suburb,
		Postcode:  req.Postcode,
		City:      req.City,
		StateCode: req.State_code,
		Country:   req.Country,
		Lat:       req.Lat,
		Lng:       req.Lng,
		Formatted: req.Formatted,
	}

	res, err := l.svcCtx.BAddressModel.Insert(l.ctx, &newItem)
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
