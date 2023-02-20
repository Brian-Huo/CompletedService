package property

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BPropertyModel = (*customBPropertyModel)(nil)

type (
	// BPropertyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBPropertyModel.
	BPropertyModel interface {
		bPropertyModel
	}

	customBPropertyModel struct {
		*defaultBPropertyModel
	}
)

// NewBPropertyModel returns a model for the database table.
func NewBPropertyModel(conn sqlx.SqlConn, c cache.CacheConf) BPropertyModel {
	return &customBPropertyModel{
		defaultBPropertyModel: newBPropertyModel(conn, c),
	}
}
