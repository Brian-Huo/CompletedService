package logic

import (
	"context"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/contractor"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveContractorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveContractorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveContractorLogic {
	return &RemoveContractorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveContractorLogic) RemoveContractor(req *types.RemoveContractorRequest) (resp *types.RemoveContractorResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	}

	if role != variables.Company {
		return nil, errorx.NewCodeError(401, "Invalid, Unauthorised action.")
	}

	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, req.Contractor_id)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Contractor not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	if uid != contractor_item.FinanceId {
		return nil, errorx.NewCodeError(404, "Invalid, Contractor not found.")
	}

	contractor_item.WorkStatus = contractor.Resigned

	err = l.svcCtx.BContractorModel.Update(l.ctx, contractor_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	return &types.RemoveContractorResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
