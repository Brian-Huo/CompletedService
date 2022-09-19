// Code generated by goctl. DO NOT EDIT!

package property

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
)

var (
	bPropertyFieldNames          = builder.RawFieldNames(&BProperty{})
	bPropertyRows                = strings.Join(bPropertyFieldNames, ",")
	bPropertyRowsExpectAutoSet   = strings.Join(stringx.Remove(bPropertyFieldNames, "`property_id`", "`create_time`", "`update_time`"), ",")
	bPropertyRowsWithPlaceHolder = strings.Join(stringx.Remove(bPropertyFieldNames, "`property_id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBPropertyPropertyIdPrefix   = "cache:bProperty:propertyId:"
	cacheBPropertyPropertyNamePrefix = "cache:bProperty:propertyName:"
)

type (
	bPropertyModel interface {
		Insert(ctx context.Context, data *BProperty) (sql.Result, error)
		FindOne(ctx context.Context, propertyId int64) (*BProperty, error)
		FindOneByPropertyName(ctx context.Context, propertyName string) (*BProperty, error)
		Update(ctx context.Context, data *BProperty) error
		Delete(ctx context.Context, propertyId int64) error
	}

	defaultBPropertyModel struct {
		sqlc.CachedConn
		table string
	}

	BProperty struct {
		PropertyId          int64  `db:"property_id"`
		PropertyName        string `db:"property_name"`
		PropertyDescription string `db:"property_description"`
		ChargeType          int64  `db:"charge_type"`
		ChargeAmount        int64  `db:"charge_amount"`
		ServiceStatus       int64  `db:"service_status"`
	}
)

func newBPropertyModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBPropertyModel {
	return &defaultBPropertyModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`b_property`",
	}
}

func (m *defaultBPropertyModel) Insert(ctx context.Context, data *BProperty) (sql.Result, error) {
	bPropertyPropertyNameKey := fmt.Sprintf("%s%v", cacheBPropertyPropertyNamePrefix, data.PropertyName)
	bPropertyPropertyIdKey := fmt.Sprintf("%s%v", cacheBPropertyPropertyIdPrefix, data.PropertyId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, bPropertyRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.PropertyName, data.PropertyDescription, data.ChargeType, data.ChargeAmount, data.ServiceStatus)
	}, bPropertyPropertyIdKey, bPropertyPropertyNameKey)
	return ret, err
}

func (m *defaultBPropertyModel) FindOne(ctx context.Context, propertyId int64) (*BProperty, error) {
	bPropertyPropertyIdKey := fmt.Sprintf("%s%v", cacheBPropertyPropertyIdPrefix, propertyId)
	var resp BProperty
	err := m.QueryRowCtx(ctx, &resp, bPropertyPropertyIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `property_id` = ? limit 1", bPropertyRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, propertyId)
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

func (m *defaultBPropertyModel) FindOneByPropertyName(ctx context.Context, propertyName string) (*BProperty, error) {
	bPropertyPropertyNameKey := fmt.Sprintf("%s%v", cacheBPropertyPropertyNamePrefix, propertyName)
	var resp BProperty
	err := m.QueryRowIndexCtx(ctx, &resp, bPropertyPropertyNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `property_name` = ? limit 1", bPropertyRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, propertyName); err != nil {
			return nil, err
		}
		return resp.PropertyId, nil
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

func (m *defaultBPropertyModel) Update(ctx context.Context, data *BProperty) error {
	bPropertyPropertyIdKey := fmt.Sprintf("%s%v", cacheBPropertyPropertyIdPrefix, data.PropertyId)
	bPropertyPropertyNameKey := fmt.Sprintf("%s%v", cacheBPropertyPropertyNamePrefix, data.PropertyName)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `property_id` = ?", m.table, bPropertyRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.PropertyName, data.PropertyDescription, data.ChargeType, data.ChargeAmount, data.ServiceStatus, data.PropertyId)
	}, bPropertyPropertyNameKey, bPropertyPropertyIdKey)
	return err
}

func (m *defaultBPropertyModel) Delete(ctx context.Context, propertyId int64) error {
	data, err := m.FindOne(ctx, propertyId)
	if err != nil {
		return err
	}

	bPropertyPropertyIdKey := fmt.Sprintf("%s%v", cacheBPropertyPropertyIdPrefix, propertyId)
	bPropertyPropertyNameKey := fmt.Sprintf("%s%v", cacheBPropertyPropertyNamePrefix, data.PropertyName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `property_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, propertyId)
	}, bPropertyPropertyIdKey, bPropertyPropertyNameKey)
	return err
}

func (m *defaultBPropertyModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBPropertyPropertyIdPrefix, primary)
}

func (m *defaultBPropertyModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `property_id` = ? limit 1", bPropertyRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBPropertyModel) tableName() string {
	return m.table
}
