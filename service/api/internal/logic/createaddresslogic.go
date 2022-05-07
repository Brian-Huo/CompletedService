package logic

import (
	"context"
	"database/sql"
	"time"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/customeraddress"

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
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)
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

	if role == variables.Customer {
		_, err := l.svcCtx.BCustomerModel.FindOne(l.ctx, uid)
		if err != nil {
			return nil, status.Error(404, "Invalid, Customer not found.")
		}

		newCustomerAddress := customeraddress.RCustomerAddress{
			CustomerId: uid,
			AddressId:  newId,
			UpdateDate: time.Now(),
		}

		_, err = l.svcCtx.RCustomerAddressModel.Insert(l.ctx, &newCustomerAddress)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	} else if role == variables.Company {
		company, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		company.RegisteredAddress = sql.NullInt64{newId, true}

		err = l.svcCtx.BCompanyModel.Update(l.ctx, company)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	}

	return &types.CreateAddressResponse{
		Address_id: newId,
	}, nil
}
