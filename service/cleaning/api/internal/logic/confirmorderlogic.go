package logic

import (
	"context"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ConfirmOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfirmOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmOrderLogic {
	return &ConfirmOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfirmOrderLogic) ConfirmOrder(req *types.ConfirmOrderRequest) (resp *types.ConfirmOrderResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, status.Error(500, "Invalid, JWT format error")
	}

	if role != variables.Company {
		return nil, status.Error(401, "Invalid, Unauthorised action.")
	}

	for _, order_id := range req.Order_list {
		err = l.svcCtx.BOrderModel.FinishStatus(l.ctx, order_id)
		if err != nil {
			logx.Info("Confirm order", order_id, "failed by finance", uid)
		}
	}

	return &types.ConfirmOrderResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
