package subscription

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RSubscriptionModel = (*customRSubscriptionModel)(nil)

var cacheRSubscriptionContractorIdPrefix = "cache:rSubscription:contractorId:"

type (
	// RSubscriptionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRSubscriptionModel.
	RSubscriptionModel interface {
		rSubscriptionModel
		DeleteByCategoryIdContractorId(ctx context.Context, categoryId int64, contractorId int64) error
		FindAllByCategory(ctx context.Context, categoryId int64) (*[]int64, error)
		FindAllByContractor(ctx context.Context, contractorId int64) (*[]int64, error)
		JoinSubscribeGroup(ctx context.Context, categoryIds *[]int64, contractorId int64) error
		LeaveSubscribeGroup(ctx context.Context, categoryIds *[]int64, contractorId int64) error
		ListSubscribeGroup(ctx context.Context, contractorId int64) (*[]int64, error)
	}

	customRSubscriptionModel struct {
		*defaultRSubscriptionModel
	}
)

// NewRSubscriptionModel returns a model for the database table.
func NewRSubscriptionModel(conn sqlx.SqlConn, c cache.CacheConf, r redis.RedisConf) RSubscriptionModel {
	return &customRSubscriptionModel{
		defaultRSubscriptionModel: newRSubscriptionModel(conn, c, r),
	}
}

func (m *defaultRSubscriptionModel) DeleteByCategoryIdContractorId(ctx context.Context, categoryId int64, contractorId int64) error {
	data, err := m.FindOneByCategoryIdContractorId(ctx, categoryId, contractorId)
	if err != nil {
		return err
	}

	rSubscriptionSubscriptionIdKey := fmt.Sprintf("%s%v", cacheRSubscriptionSubscriptionIdPrefix, data.SubscriptionId)
	rSubscriptionCategoryIdContractorIdKey := fmt.Sprintf("%s%v:%v", cacheRSubscriptionCategoryIdContractorIdPrefix, data.CategoryId, data.ContractorId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `subscription_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, data.SubscriptionId)
	}, rSubscriptionSubscriptionIdKey, rSubscriptionCategoryIdContractorIdKey)
	return err
}

func (m *defaultRSubscriptionModel) FindAllByCategory(ctx context.Context, categoryId int64) (*[]int64, error) {
	var resp []int64
	query := fmt.Sprintf("select `contractor_id` from %s where `category_id` = ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, categoryId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRSubscriptionModel) FindAllByContractor(ctx context.Context, contractorId int64) (*[]int64, error) {
	var resp []int64
	query := fmt.Sprintf("select `category_id` from %s where `contractor_id` = ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, contractorId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRSubscriptionModel) JoinSubscribeGroup(ctx context.Context, categoryIds *[]int64, contractorId int64) error {
	rSubscriptionContractorIdKey := fmt.Sprintf("%s%v", cacheRSubscriptionContractorIdPrefix, contractorId)
	for _, categoryId := range *categoryIds {
		_, err := m.Insert(ctx, &RSubscription{CategoryId: categoryId, ContractorId: categoryId})
		if err != nil {
			return err
		}
	}
	_, err := m.Sadd(rSubscriptionContractorIdKey, categoryIds)
	return err
}

func (m *defaultRSubscriptionModel) LeaveSubscribeGroup(ctx context.Context, categoryIds *[]int64, contractorId int64) error {
	rSubscriptionContractorIdKey := fmt.Sprintf("%s%v", cacheRSubscriptionContractorIdPrefix, contractorId)
	for _, categoryId := range *categoryIds {
		err := m.DeleteByCategoryIdContractorId(ctx, categoryId, categoryId)
		if err != nil {
			return err
		}
	}
	_, err := m.Srem(rSubscriptionContractorIdKey, categoryIds)
	return err
}

func (m *defaultRSubscriptionModel) ListSubscribeGroup(ctx context.Context, contractorId int64) (*[]int64, error) {
	rSubscriptionContractorIdKey := fmt.Sprintf("%s%v", cacheRSubscriptionContractorIdPrefix, contractorId)
	ret_str, err := m.Redis.Smembers(rSubscriptionContractorIdKey)
	if err != nil {
		return nil, err
	}

	if len(ret_str) <= 0 {
		ret, err := m.FindAllByContractor(ctx, contractorId)
		if err != nil {
			return nil, err
		}
		go m.Sadd(rSubscriptionContractorIdKey, ret)
		return ret, nil
	}

	var resp []int64
	for _, val := range ret_str {
		category_id, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			resp = append(resp, category_id)
		}
	}
	return &resp, nil
}
