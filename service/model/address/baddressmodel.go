package address

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BAddressModel = (*customBAddressModel)(nil)

type (
	// BAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBAddressModel.
	BAddressModel interface {
		bAddressModel
	}

	customBAddressModel struct {
		*defaultBAddressModel
	}
)

// NewBAddressModel returns a model for the database table.
func NewBAddressModel(conn sqlx.SqlConn, c cache.CacheConf) BAddressModel {
	return &customBAddressModel{
		defaultBAddressModel: newBAddressModel(conn, c),
	}
}
