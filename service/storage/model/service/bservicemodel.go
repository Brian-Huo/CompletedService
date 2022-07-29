package service

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BServiceModel = (*customBServiceModel)(nil)

type (
	// BServiceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBServiceModel.
	BServiceModel interface {
		bServiceModel
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
