package orderdelay

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var _ ROrderDelayModel = (*customROrderDelayModel)(nil)

type (
	// ROrderDelayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customROrderDelayModel.
	ROrderDelayModel interface {
		rOrderDelayModel
	}

	customROrderDelayModel struct {
		*defaultROrderDelayModel
	}
)

// NewROrderDelayModel returns a model for the database table.
func NewROrderDelayModel(c redis.RedisConf) ROrderDelayModel {
	return &customROrderDelayModel{
		defaultROrderDelayModel: newROrderDelayModel(c),
	}
}
