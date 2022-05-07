package logic

import (
	"context"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/address"

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
	res, err := l.svcCtx.BAddressModel.FindOne(l.ctx, req.Address_id)
	if err != nil {
		if err == address.ErrNotFound {
			return nil, status.Error(404, "Invalid, Address not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &types.DetailAddressResponse{
		Address_id:      res.AddressId,
		Address_details: res.StateCode,
		Suburb:          res.Suburb,
		Postcode:        res.Postcode,
		State_code:      res.StateCode,
		Country:         res.Country.String,
	}, nil
}
