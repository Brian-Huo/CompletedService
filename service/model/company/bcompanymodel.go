package company

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BCompanyModel = (*customBCompanyModel)(nil)

type (
	// BCompanyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBCompanyModel.
	BCompanyModel interface {
		bCompanyModel
	}

	customBCompanyModel struct {
		*defaultBCompanyModel
	}
)

// NewBCompanyModel returns a model for the database table.
func NewBCompanyModel(conn sqlx.SqlConn, c cache.CacheConf) BCompanyModel {
	return &customBCompanyModel{
		defaultBCompanyModel: newBCompanyModel(conn, c),
	}
}
