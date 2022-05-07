package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/address"

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
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
	if role == variables.Employee {
		return nil, status.Error(401, "Invalid, Not customer/company.")
	}

	// check address id vaild for customer or compay
	if role == variables.Company {
		res, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		if res.RegisteredAddress.Valid && res.RegisteredAddress.Int64 != req.Address_id {
			return nil, status.Error(401, "Invalid company address id.")
		}
	} else if role == variables.Customer {
		_, err := l.svcCtx.RCustomerAddressModel.FindOne(l.ctx, uid, req.Address_id)
		if err != nil {
			return nil, status.Error(401, "Invalid customer address id.")
		}
	}

	err = l.svcCtx.BAddressModel.Update(l.ctx, &address.BAddress{
		AddressDetails: req.Address_details,
		Suburb:         req.Suburb,
		Postcode:       req.Postcode,
		StateCode:      req.State_code,
		Country:        sql.NullString{req.Country, req.Country != ""},
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.UpdateAddressResponse{}, nil
}
