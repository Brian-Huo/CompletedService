package logic

import (
	"context"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/contractor"
	"cleaningservice/service/model/subscribegroup"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, status.Error(401, "Invalid, Not contractor.")
	}

	// Get contractor detail
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, status.Error(404, "Invalid, Contractor not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Get address details
	address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, contractor_item.AddressId.Int64)
	if err != nil {
		if err == address.ErrNotFound {
			return nil, errorx.NewCodeError(401, "Invalid, Please update your address details first.")
		}
		return nil, status.Error(500, err.Error())
	}

	// Unsubscribe
	subscribegroup_item, err := l.svcCtx.BSubscribeGroupModel.FindOneByCategoryLocation(l.ctx, req.Category_id, address_item.City)
	if err != nil {
		if err == subscribegroup.ErrNotFound {
			return nil, status.Error(404, "Invalid, Subscribe group not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.RSubscribeRecordModel.Delete(l.ctx, subscribegroup_item.GroupId, uid)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	_, err = l.svcCtx.BSubscriptionModel.Delete(subscribegroup_item.GroupId, uid)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.LeaveSubscribeGroupResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
