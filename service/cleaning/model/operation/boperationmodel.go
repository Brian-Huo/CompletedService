package operation

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BOperationModel = (*customBOperationModel)(nil)

type (
	// BOperationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBOperationModel.
	BOperationModel interface {
		bOperationModel
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
