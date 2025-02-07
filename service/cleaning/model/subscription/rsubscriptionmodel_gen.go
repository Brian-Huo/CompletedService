// Code generated by goctl. DO NOT EDIT!

package subscription

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var (
	rSubscriptionFieldNames          = builder.RawFieldNames(&RSubscription{})
	rSubscriptionRows                = strings.Join(rSubscriptionFieldNames, ",")
	rSubscriptionRowsExpectAutoSet   = strings.Join(stringx.Remove(rSubscriptionFieldNames, "`subscription_id`", "`create_time`", "`update_time`"), ",")
	rSubscriptionRowsWithPlaceHolder = strings.Join(stringx.Remove(rSubscriptionFieldNames, "`subscription_id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheRSubscriptionSubscriptionIdPrefix         = "cache:rSubscription:subscriptionId:"
	cacheRSubscriptionCategoryIdContractorIdPrefix = "cache:rSubscription:categoryId:contractorId:"
)

type (
	rSubscriptionModel interface {
		Insert(ctx context.Context, data *RSubscription) (sql.Result, error)
		FindOne(ctx context.Context, subscriptionId int64) (*RSubscription, error)
		FindOneByCategoryIdContractorId(ctx context.Context, categoryId int64, contractorId int64) (*RSubscription, error)
		Update(ctx context.Context, data *RSubscription) error
		Delete(ctx context.Context, subscriptionId int64) error
	}

	defaultRSubscriptionModel struct {
		sqlc.CachedConn
		redis.Redis
		table string
	}

	RSubscription struct {
		SubscriptionId int64 `db:"subscription_id"`
		CategoryId     int64 `db:"category_id"`
		ContractorId   int64 `db:"contractor_id"`
	}
)

func newRSubscriptionModel(conn sqlx.SqlConn, c cache.CacheConf, r redis.RedisConf) *defaultRSubscriptionModel {
	return &defaultRSubscriptionModel{
		CachedConn: sqlc.NewConn(conn, c),
		Redis:  	*r.NewRedis(),
		table:      "`r_subscription`",
	}
}

func (m *defaultRSubscriptionModel) Insert(ctx context.Context, data *RSubscription) (sql.Result, error) {
	rSubscriptionSubscriptionIdKey := fmt.Sprintf("%s%v", cacheRSubscriptionSubscriptionIdPrefix, data.SubscriptionId)
	rSubscriptionCategoryIdContractorIdKey := fmt.Sprintf("%s%v:%v", cacheRSubscriptionCategoryIdContractorIdPrefix, data.CategoryId, data.ContractorId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, rSubscriptionRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CategoryId, data.ContractorId)
	}, rSubscriptionSubscriptionIdKey, rSubscriptionCategoryIdContractorIdKey)
	return ret, err
}

func (m *defaultRSubscriptionModel) FindOne(ctx context.Context, subscriptionId int64) (*RSubscription, error) {
	rSubscriptionSubscriptionIdKey := fmt.Sprintf("%s%v", cacheRSubscriptionSubscriptionIdPrefix, subscriptionId)
	var resp RSubscription
	err := m.QueryRowCtx(ctx, &resp, rSubscriptionSubscriptionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `subscription_id` = ? limit 1", rSubscriptionRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, subscriptionId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRSubscriptionModel) FindOneByCategoryIdContractorId(ctx context.Context, categoryId int64, contractorId int64) (*RSubscription, error) {
	rSubscriptionCategoryIdContractorIdKey := fmt.Sprintf("%s%v:%v", cacheRSubscriptionCategoryIdContractorIdPrefix, categoryId, contractorId)
	var resp RSubscription
	err := m.QueryRowIndexCtx(ctx, &resp, rSubscriptionCategoryIdContractorIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `category_id` = ? and `contractor_id` = ? limit 1", rSubscriptionRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, categoryId, contractorId); err != nil {
			return nil, err
		}
		return resp.SubscriptionId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRSubscriptionModel) Update(ctx context.Context, data *RSubscription) error {
	rSubscriptionSubscriptionIdKey := fmt.Sprintf("%s%v", cacheRSubscriptionSubscriptionIdPrefix, data.SubscriptionId)
	rSubscriptionCategoryIdContractorIdKey := fmt.Sprintf("%s%v:%v", cacheRSubscriptionCategoryIdContractorIdPrefix, data.CategoryId, data.ContractorId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `subscription_id` = ?", m.table, rSubscriptionRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CategoryId, data.ContractorId, data.SubscriptionId)
	}, rSubscriptionSubscriptionIdKey, rSubscriptionCategoryIdContractorIdKey)
	return err
}

func (m *defaultRSubscriptionModel) Delete(ctx context.Context, subscriptionId int64) error {
	data, err := m.FindOne(ctx, subscriptionId)
	if err != nil {
		return err
	}

	rSubscriptionSubscriptionIdKey := fmt.Sprintf("%s%v", cacheRSubscriptionSubscriptionIdPrefix, subscriptionId)
	rSubscriptionCategoryIdContractorIdKey := fmt.Sprintf("%s%v:%v", cacheRSubscriptionCategoryIdContractorIdPrefix, data.CategoryId, data.ContractorId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `subscription_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, subscriptionId)
	}, rSubscriptionSubscriptionIdKey, rSubscriptionCategoryIdContractorIdKey)
	return err
}

func (m *defaultRSubscriptionModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheRSubscriptionSubscriptionIdPrefix, primary)
}

func (m *defaultRSubscriptionModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `subscription_id` = ? limit 1", rSubscriptionRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRSubscriptionModel) tableName() string {
	return m.table
}
