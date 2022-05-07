package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/company"
	"cleaningservice/service/model/customeraddress"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemoveAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAddressLogic {
	return &RemoveAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveAddressLogic) RemoveAddress(req *types.RemoveAddressRequest) (resp *types.RemoveAddressResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)

	if role == variables.Employee {
		return nil, status.Error(401, "Invalid, Unauthorised action.")
	} else if role == variables.Company {
		comp, err := l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
		if err != nil {
			if err == company.ErrNotFound {
				return nil, status.Error(404, "Invalid, Company not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		comp.RegisteredAddress = sql.NullInt64{0, false}

		err = l.svcCtx.BCompanyModel.Update(l.ctx, comp)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	} else if role == variables.Customer {
		err := l.svcCtx.RCustomerAddressModel.Delete(l.ctx, uid, req.Address_id)
		if err != nil {
			if err == customeraddress.ErrNotFound {
				return nil, status.Error(404, "Invalid, Customer address record not found.")
			}
			return nil, status.Error(500, err.Error())
		}
	}

	go l.svcCtx.BAddressModel.Delete(l.ctx, req.Address_id)

	return
}
