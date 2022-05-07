package customer

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BCustomerModel = (*customBCustomerModel)(nil)

type (
	// BCustomerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBCustomerModel.
	BCustomerModel interface {
		bCustomerModel
	}

	customBCustomerModel struct {
		*defaultBCustomerModel
	}
)

// NewBCustomerModel returns a model for the database table.
func NewBCustomerModel(conn sqlx.SqlConn, c cache.CacheConf) BCustomerModel {
	return &customBCustomerModel{
		defaultBCustomerModel: newBCustomerModel(conn, c),
	}
}
