package operation

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BOperationModel = (*customBOperationModel)(nil)

type (
	// BOperationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBOperationModel.
	BOperationModel interface {
		bOperationModel
		FindAllByContractor(ctx context.Context, contractorId int64) ([]*BOperation, error)
		FindAllByOrder(ctx context.Context, orderId int64) ([]*BOperation, error)
		DeleteAllByContractor(ctx context.Context, contractorId int64) error
		DeleteAllByOrder(ctx context.Context, orderId int64) error
		RecordAccept(ctx context.Context, contractorId int64, orderId int64) (sql.Result, error)
		RecordDecline(ctx context.Context, contractorId int64, orderId int64) (sql.Result, error)
		RecordTransfer(ctx context.Context, contractorId int64, orderId int64) (sql.Result, error)
	}

	customBOperationModel struct {
		*defaultBOperationModel
	}
)

// NewBOperationModel returns a model for the database table.
func NewBOperationModel(conn sqlx.SqlConn, c cache.CacheConf) BOperationModel {
	return &customBOperationModel{
		defaultBOperationModel: newBOperationModel(conn, c),
	}
}

func (m *defaultBOperationModel) FindAllByContractor(ctx context.Context, contractorId int64) ([]*BOperation, error) {
	var resp []*BOperation
	query := fmt.Sprintf("select %s from %s where `contractor_id` = ?", bOperationRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, contractorId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBOperationModel) FindAllByOrder(ctx context.Context, orderId int64) ([]*BOperation, error) {
	var resp []*BOperation
	query := fmt.Sprintf("select %s from %s where `order_id` = ?", bOperationRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, orderId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBOperationModel) DeleteAllByContractor(ctx context.Context, contractorId int64) error {
	query := fmt.Sprintf("delete from %s where `contractor_id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, contractorId)
	return err
}

func (m *defaultBOperationModel) DeleteAllByOrder(ctx context.Context, orderId int64) error {
	query := fmt.Sprintf("delete from %s where `order_id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, orderId)
	return err
}

func (m *defaultBOperationModel) RecordAccept(ctx context.Context, contractorId int64, orderId int64) (sql.Result, error) {
	bOperationOperationIdKey := fmt.Sprintf("%s%v", cacheBOperationOperationIdPrefix, nil)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, bOperationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, contractorId, orderId, Accept)
	}, bOperationOperationIdKey)
	return ret, err
}

func (m *defaultBOperationModel) RecordDecline(ctx context.Context, contractorId int64, orderId int64) (sql.Result, error) {
	bOperationOperationIdKey := fmt.Sprintf("%s%v", cacheBOperationOperationIdPrefix, nil)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, bOperationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, contractorId, orderId, Decline)
	}, bOperationOperationIdKey)
	return ret, err
}

func (m *defaultBOperationModel) RecordTransfer(ctx context.Context, contractorId int64, orderId int64) (sql.Result, error) {
	bOperationOperationIdKey := fmt.Sprintf("%s%v", cacheBOperationOperationIdPrefix, nil)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, bOperationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, contractorId, orderId, Transfer)
	}, bOperationOperationIdKey)
	return ret, err
}
