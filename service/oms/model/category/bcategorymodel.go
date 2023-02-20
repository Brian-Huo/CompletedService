package category

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BCategoryModel = (*customBCategoryModel)(nil)

type (
	// BCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBCategoryModel.
	BCategoryModel interface {
		bCategoryModel
		List(ctx context.Context) ([]*BCategory, error)
	}

	customBCategoryModel struct {
		*defaultBCategoryModel
	}
)

// NewBCategoryModel returns a model for the database table.
func NewBCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) BCategoryModel {
	return &customBCategoryModel{
		defaultBCategoryModel: newBCategoryModel(conn, c),
	}
}

func (m *defaultBCategoryModel) List(ctx context.Context) ([]*BCategory, error) {
	var resp []*BCategory

	query := fmt.Sprintf("select %s from %s", bCategoryRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
