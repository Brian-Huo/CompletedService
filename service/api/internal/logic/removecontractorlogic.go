package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/contractor"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	if role != variables.Company {
		return nil, status.Error(401, "Invalid, Unauthorised action.")
	}

	cont, err := l.svcCtx.BContractorModel.FindOne(l.ctx, req.Contractor_id)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, status.Error(404, "Invalid, Contractor not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	if uid != cont.FinanceId {
		return nil, status.Error(404, "Invalid, Contractor not found.")
	}

	cont.WorkStatus = int64(variables.Resigned)

	err = l.svcCtx.BContractorModel.Update(l.ctx, cont)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.RemoveContractorResponse{}, nil
}
