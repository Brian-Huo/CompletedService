package payment

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BPaymentModel = (*customBPaymentModel)(nil)

type (
	// BPaymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBPaymentModel.
	BPaymentModel interface {
		bPaymentModel
	}

	customBPaymentModel struct {
		*defaultBPaymentModel
	}
)

// NewBPaymentModel returns a model for the database table.
func NewBPaymentModel(conn sqlx.SqlConn, c cache.CacheConf) BPaymentModel {
	return &customBPaymentModel{
		defaultBPaymentModel: newBPaymentModel(conn, c),
	}
}
