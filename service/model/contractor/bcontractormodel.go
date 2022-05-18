package contractor

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BContractorModel = (*customBContractorModel)(nil)

type (
	// BContractorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBContractorModel.
	BContractorModel interface {
		bContractorModel
	}

	customBContractorModel struct {
		*defaultBContractorModel
	}
)

// NewBContractorModel returns a model for the database table.
func NewBContractorModel(conn sqlx.SqlConn, c cache.CacheConf) BContractorModel {
	return &customBContractorModel{
		defaultBContractorModel: newBContractorModel(conn, c),
	}
}
