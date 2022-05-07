package logic

import (
	"context"
	"database/sql"
	"log"

	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/company"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCompanyLogic {
	return &CreateCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCompanyLogic) CreateCompany(req *types.CreateCompanyRequest) (resp *types.CreateCompanyResponse, err error) {
	uid := l.ctx.Value("uid").(int64)
	role := l.ctx.Value("role").(int)

	if role != 100 && uid != 0 {
		log.Println("Backend broken, security leak...")
		return nil, status.Error(500, err.Error())
	}

	newItem := company.BCompany{
		CompanyName:       req.Company_name,
		PaymentId:         sql.NullInt64{req.Payment_id, req.Payment_id != 0},
		DirectorName:      sql.NullString{req.Director_name, req.Director_name != ""},
		ContactDetails:    req.Contact_details,
		RegisteredAddress: sql.NullInt64{req.Registered_address, req.Registered_address != 0},
		DepositeRate:      int64(req.Deposite_rate),
	}

	res, err := l.svcCtx.BCompanyModel.Insert(l.ctx, &newItem)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateCompanyResponse{
		Company_id: newId,
	}, nil
}
