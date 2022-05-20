package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/operation"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DetailOperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailOperationLogic {
	return &DetailOperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailOperationLogic) DetailOperation(req *types.DetailOperationRequest) (resp *types.DetailOperationResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role == variables.Company {
		res, err := l.svcCtx.BOperationModel.FindOne(l.ctx, req.Operation_id)
		if err != nil {
			if err == operation.ErrNotFound {
				return nil, status.Error(404, "Invalid, Operation record not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		cont, err := l.svcCtx.BContractorModel.FindOne(l.ctx, res.ContractorId)
		if err != nil || uid != cont.FinanceId {
			if err == operation.ErrNotFound {
				return nil, status.Error(404, "Invalid, Contractor not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		return &types.DetailOperationResponse{
			Operation_id:  res.OperationId,
			Contractor_id: res.ContractorId,
			Order_id:      res.OrderId,
			Operation:     res.Operation,
			Issue_date:    resp.Issue_date,
		}, nil
	}

	return nil, status.Error(404, "Invalid, Operation record not found.")
}
