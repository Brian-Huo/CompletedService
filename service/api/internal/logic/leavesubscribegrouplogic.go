package logic

import (
	"context"
	"database/sql"

	"cleaningservice/common/jwtx"
	"cleaningservice/common/variables"
	"cleaningservice/service/api/internal/svc"
	"cleaningservice/service/api/internal/types"
	"cleaningservice/service/model/contractor"
	"cleaningservice/util"

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

	// Update category list in contractor
	previous_category := util.StringToIntArray(contractor_item.CategoryList.String)
	new_category := util.IntArrayToString(util.DisjointIntArray(previous_category, req.Category_list))
	contractor_item.CategoryList = sql.NullString{new_category, new_category != ""}

	// Unsubscribe all category group
	for _, category_id := range req.Category_list {
		_, err := l.svcCtx.BCategoryModel.FindOne(l.ctx, category_id)
		if err != nil {
			continue
		}

		err = l.svcCtx.RSubscribeRecordModel.Delete(l.ctx, category_id, uid)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		_, err = l.svcCtx.BSubscriptionModel.Delete(category_id, uid)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
	}

	err = l.svcCtx.BContractorModel.Update(l.ctx, contractor_item)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.LeaveSubscribeGroupResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
