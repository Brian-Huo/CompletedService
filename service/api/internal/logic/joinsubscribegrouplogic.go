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
	"cleaningservice/service/model/subscriberecord"
	"cleaningservice/service/model/subscription"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type JoinSubscribeGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinSubscribeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinSubscribeGroupLogic {
	return &JoinSubscribeGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinSubscribeGroupLogic) JoinSubscribeGroup(req *types.JoinSubscribeGroupRequest) (resp *types.JoinSubscribeGroupResponse, err error) {
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

	// Subscribe all category groups
	for _, category_id := range req.Category_list {
		subscribegroup_item, err := l.svcCtx.BSubscribeGroupModel.FindOneByCategoryLocation(l.ctx, category_id, address_item.City)
		if err != nil {
			if err == subscribegroup.ErrNotFound {
				return nil, status.Error(404, "Invalid, Subscribe group not found.")
			}
			return nil, status.Error(500, err.Error())
		}

		_, err = l.svcCtx.RSubscribeRecordModel.Insert(l.ctx, &subscriberecord.RSubscribeRecord{
			GroupId:      subscribegroup_item.GroupId,
			ContractorId: uid,
		})
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		_, err = l.svcCtx.BSubscriptionModel.Insert(&subscription.BSubscription{
			GroupId:      subscribegroup_item.GroupId,
			ContractorId: uid,
		})
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	}

	return &types.JoinSubscribeGroupResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
