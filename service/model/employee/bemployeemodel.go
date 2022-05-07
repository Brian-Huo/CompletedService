package employee

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BEmployeeModel = (*customBEmployeeModel)(nil)

type (
	// BEmployeeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBEmployeeModel.
	BEmployeeModel interface {
		bEmployeeModel
	}

	customBEmployeeModel struct {
		*defaultBEmployeeModel
	}
)

// NewBEmployeeModel returns a model for the database table.
func NewBEmployeeModel(conn sqlx.SqlConn, c cache.CacheConf) BEmployeeModel {
	return &customBEmployeeModel{
		defaultBEmployeeModel: newBEmployeeModel(conn, c),
	}
}
