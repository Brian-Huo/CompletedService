package order

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BOrderModel = (*customBOrderModel)(nil)

type (
	// BOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBOrderModel.
	BOrderModel interface {
		bOrderModel
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
