package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/address"

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
	_, role, _ := jwtx.GetTokenDetails(l.ctx)
	if role == variables.Employee {
		return nil, status.Error(401, "Invalid, Not customer/company.")
	}

	newItem := address.BAddress{
		AddressDetails: req.Address_details,
		Suburb:         req.Suburb,
		Postcode:       req.Postcode,
		StateCode:      req.State_code,
		Country:        sql.NullString{req.Country, req.Country != ""},
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
