package order

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BOrderModel = (*customBOrderModel)(nil)

type (
	// BOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBOrderModel.
	BOrderModel interface {
		bOrderModel
		FindAllByAddress(ctx context.Context, addressId int64) ([]*BOrder, error)
		FindAllByFinance(ctx context.Context, financeId int64) ([]*BOrder, error)
		FindAllByCustomer(ctx context.Context, customerId int64) ([]*BOrder, error)
		FindAllByContractor(ctx context.Context, contractorId int64) ([]*BOrder, error)
	}

	customBOrderModel struct {
		*defaultBOrderModel
	}
)

// NewBOrderModel returns a model for the database table.
func NewBOrderModel(conn sqlx.SqlConn, c cache.CacheConf) BOrderModel {
	return &customBOrderModel{
		defaultBOrderModel: newBOrderModel(conn, c),
	}
}

func (m *defaultBOrderModel) FindAllByAddress(ctx context.Context, addressId int64) ([]*BOrder, error) {
	var resp []*BOrder
	query := fmt.Sprintf("select %s from %s where `address_id` = ?", bOrderRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, addressId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBOrderModel) FindAllByFinance(ctx context.Context, financeId int64) ([]*BOrder, error) {
	var resp []*BOrder
	query := fmt.Sprintf("select %s from %s where `finance_id` = ?", bOrderRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, financeId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBOrderModel) FindAllByCustomer(ctx context.Context, customerId int64) ([]*BOrder, error) {
	var resp []*BOrder
	query := fmt.Sprintf("select %s from %s where `customer_id` = ?", bOrderRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, customerId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBOrderModel) FindAllByContractor(ctx context.Context, contractorId int64) ([]*BOrder, error) {
	var resp []*BOrder
	query := fmt.Sprintf("select %s from %s where `contractor_id` = ?", bOrderRows, m.table)
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
