package subscribegroup

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BSubscribeGroupModel = (*customBSubscribeGroupModel)(nil)

type (
	// BSubscribeGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBSubscribeGroupModel.
	BSubscribeGroupModel interface {
		bSubscribeGroupModel
	}

	customBSubscribeGroupModel struct {
		*defaultBSubscribeGroupModel
	}
)

// NewBSubscribeGroupModel returns a model for the database table.
func NewBSubscribeGroupModel(conn sqlx.SqlConn, c cache.CacheConf) BSubscribeGroupModel {
	return &customBSubscribeGroupModel{
		defaultBSubscribeGroupModel: newBSubscribeGroupModel(conn, c),
	}
}
