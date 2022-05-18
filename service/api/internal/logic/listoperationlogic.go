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

type ListOperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOperationLogic {
	return &ListOperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListOperationLogic) ListOperation(req *types.ListOperationRequest) (resp *types.ListOperationResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role != variables.Employee {
		return nil, status.Error(401, "Invalid, Not employee.")
	}

	res, err := l.svcCtx.BOperationModel.FindAllByContractor(l.ctx, uid)
	if err != nil {
		if err == operation.ErrNotFound {
			return nil, status.Error(404, "Invalid, Operation not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	allItems := []types.DetailOperationResponse{}

	for _, item := range res {
		newItem := types.DetailOperationResponse{
			Operation_id:  item.OperationId,
			Contractor_id: item.ContractorId,
			Order_id:      item.OrderId,
			Operation:     item.Operation,
			Issue_date:    item.IssueDate.String(),
		}

		allItems = append(allItems, newItem)
	}

	return &types.ListOperationResponse{
		Items: allItems,
	}, nil
}
