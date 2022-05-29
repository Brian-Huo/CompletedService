package orderrecommend

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var _ ROrderRecommendModel = (*customROrderRecommendModel)(nil)

type (
	// ROrderRecommendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customROrderRecommendModel.
	ROrderRecommendModel interface {
		rOrderRecommendModel
	}

	customROrderRecommendModel struct {
		*defaultROrderRecommendModel
	}
)

// NewROrderRecommendModel returns a model for the database table.
func NewROrderRecommendModel(c redis.RedisConf) ROrderRecommendModel {
	return &customROrderRecommendModel{
		defaultROrderRecommendModel: newROrderRecommendModel(c),
	}
}
