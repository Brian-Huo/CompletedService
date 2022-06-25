package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/subscription"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type InitSystemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitSystemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitSystemLogic {
	return &InitSystemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitSystemLogic) InitSystem(req *types.InitSystemRequest) (resp *types.InitSystemResponse, err error) {
	// Valide JWT token
	_, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}
	if role != variables.Admin {
		return nil, status.Error(401, "Invalid, Not Admin.")
	}

	// Inisialize Redis database
	// Get all subscription records
	record_items, err := l.svcCtx.RSubscribeRecordModel.List(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// Insert all subscriptions
	for _, record_item := range record_items {
		_, err := l.svcCtx.BSubscriptionModel.Insert(&subscription.BSubscription{
			GroupId:      record_item.CategoryId,
			ContractorId: record_item.ContractorId,
		})
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	}

	return &types.InitSystemResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
