package broadcast

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var _ BBroadcastModel = (*customBBroadcastModel)(nil)

type (
	// BBroadcastModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBBroadcastModel.
	BBroadcastModel interface {
		bBroadcastModel
	}

	customBBroadcastModel struct {
		*defaultBBroadcastModel
	}
)

// NewBBroadcastModel returns a model for the database table.
func NewBBroadcastModel(c redis.RedisConf) BBroadcastModel {
	return &customBBroadcastModel{
		defaultBBroadcastModel: newBBroadcastModel(c),
	}
}
