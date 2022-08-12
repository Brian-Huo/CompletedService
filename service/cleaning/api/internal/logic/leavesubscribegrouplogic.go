package logic

import (
	"context"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/contractor"

	"github.com/zeromicro/go-zero/core/logx"
)

type LeaveSubscribeGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLeaveSubscribeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveSubscribeGroupLogic {
	return &LeaveSubscribeGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LeaveSubscribeGroupLogic) LeaveSubscribeGroup(req *types.LeaveSubscribeGroupRequest) (resp *types.LeaveSubscribeGroupResponse, err error) {
	uid, role, err := jwtx.GetTokenDetails(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, errorx.NewCodeError(401, "Invalid, Not contractor.")
	}

	// Get contractor detail
	_, err = l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Contractor not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Leave subscription group
	err = l.svcCtx.RSubscriptionModel.LeaveSubscribeGroup(l.ctx, &req.Category_list, uid)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	return &types.LeaveSubscribeGroupResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
