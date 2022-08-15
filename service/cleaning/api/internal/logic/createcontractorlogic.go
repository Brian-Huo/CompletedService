package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/contractor"
	"cleaningservice/util"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateContractorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContractorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContractorLogic {
	return &CreateContractorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContractorLogic) CreateContractor(req *types.CreateContractorRequest) (resp *types.CreateContractorResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		logx.Info("jwt issue")
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	if role == variables.Company {
		_, err = l.svcCtx.BCompanyModel.FindOne(l.ctx, uid)
		if err != nil {
			return nil, status.Error(404, "Invalid, Company not found.")
		}
	} else {
		return nil, status.Error(401, "Invalid, Unauthorised action.")
	}

	newItem := contractor.BContractor{
		ContractorPhoto: sql.NullString{String: req.Contractor_photo, Valid: req.Contractor_photo != ""},
		ContractorName:  req.Contractor_name,
		ContactDetails:  req.Contact_details,
		ContractorType:  contractor.Employee,
		FinanceId:       uid,
		AddressId:       sql.NullInt64{Int64: 0, Valid: false},
		LinkCode:        util.RandStringBytesMaskImprSrcUnsafe(8),
		WorkStatus:      contractor.Await,
		OrderId:         sql.NullInt64{Int64: 0, Valid: false},
	}

	res, err := l.svcCtx.BContractorModel.Insert(l.ctx, &newItem)
	if err != nil {
		logx.Info("test here")
		return nil, status.Error(500, err.Error())
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateContractorResponse{
		Contractor_id: newId,
	}, nil
}
