package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/errorx"
	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/contractor"
	"cleaningservice/service/model/subscriberecord"
	"cleaningservice/service/model/subscription"
	"cleaningservice/util"

	"github.com/zeromicro/go-zero/core/logx"
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
		return nil, errorx.NewCodeError(500, "Invalid, JWT format error")
	} else if role != variables.Contractor {
		return nil, errorx.NewCodeError(401, "Invalid, Not contractor.")
	}

	// Get contractor detail
	contractor_item, err := l.svcCtx.BContractorModel.FindOne(l.ctx, uid)
	if err != nil {
		if err == contractor.ErrNotFound {
			return nil, errorx.NewCodeError(404, "Invalid, Contractor not found.")
		}
		return nil, errorx.NewCodeError(500, err.Error())
	}

	// Update category list in contractor
	previous_category := util.StringToIntArray(contractor_item.CategoryList.String)
	new_category := util.IntArrayToString(util.UnionIntArray(previous_category, req.Category_list))
	contractor_item.CategoryList = sql.NullString{new_category, new_category != ""}

	// Subscribe all category groups
	for _, category_id := range req.Category_list {
		_, err := l.svcCtx.BCategoryModel.FindOne(l.ctx, category_id)
		if err != nil {
			continue
		}

		_, err = l.svcCtx.RSubscribeRecordModel.Insert(l.ctx, &subscriberecord.RSubscribeRecord{
			CategoryId:   category_id,
			ContractorId: uid,
		})
		if err != nil {
			return nil, errorx.NewCodeError(500, err.Error())
		}

		_, err = l.svcCtx.BSubscriptionModel.Insert(&subscription.BSubscription{
			GroupId:      category_id,
			ContractorId: uid,
		})
		if err != nil {
			return nil, errorx.NewCodeError(500, err.Error())
		}
	}

	err = l.svcCtx.BContractorModel.Update(l.ctx, contractor_item)
	if err != nil {
		return nil, errorx.NewCodeError(500, err.Error())
	}

	return &types.JoinSubscribeGroupResponse{
		Code: 200,
		Msg:  "Success",
	}, nil
}
