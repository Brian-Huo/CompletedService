package category

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BCategoryModel = (*customBCategoryModel)(nil)

type (
	// BCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBCategoryModel.
	BCategoryModel interface {
		bCategoryModel
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
