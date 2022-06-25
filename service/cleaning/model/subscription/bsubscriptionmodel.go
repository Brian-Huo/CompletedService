package subscription

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var _ BSubscriptionModel = (*customBSubscriptionModel)(nil)

type (
	// BSubscriptionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBSubscriptionModel.
	BSubscriptionModel interface {
		bSubscriptionModel
	}

	customBSubscriptionModel struct {
		*defaultBSubscriptionModel
	}
)

// NewBSubscriptionModel returns a model for the database table.
func NewBSubscriptionModel(c redis.RedisConf) BSubscriptionModel {
	return &customBSubscriptionModel{
		defaultBSubscriptionModel: newBSubscriptionModel(c),
	}
}
