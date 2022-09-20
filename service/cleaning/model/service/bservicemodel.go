package service

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BServiceModel = (*customBServiceModel)(nil)

type (
	// BServiceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBServiceModel.
	BServiceModel interface {
		bServiceModel
		FindAllByCategory(ctx context.Context, categoryId int64) ([]*BService, error)
	}

	customBServiceModel struct {
		*defaultBServiceModel
	}
)

// NewBServiceModel returns a model for the database table.
func NewBServiceModel(conn sqlx.SqlConn, c cache.CacheConf) BServiceModel {
	return &customBServiceModel{
		defaultBServiceModel: newBServiceModel(conn, c),
	}
}

func (m *defaultBServiceModel) FindAllByCategory(ctx context.Context, categoryId int64) ([]*BService, error) {
	var resp []*BService
	query := fmt.Sprintf("select %s from %s where `service_type` = ?", bServiceRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, categoryId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
