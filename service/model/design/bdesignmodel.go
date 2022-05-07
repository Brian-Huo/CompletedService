package design

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BDesignModel = (*customBDesignModel)(nil)

type (
	// BDesignModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBDesignModel.
	BDesignModel interface {
		bDesignModel
	}

	customBDesignModel struct {
		*defaultBDesignModel
	}
)

// NewBDesignModel returns a model for the database table.
func NewBDesignModel(conn sqlx.SqlConn, c cache.CacheConf) BDesignModel {
	return &customBDesignModel{
		defaultBDesignModel: newBDesignModel(conn, c),
	}
}
