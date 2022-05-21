package schedule

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var _ BScheduleModel = (*customBScheduleModel)(nil)

type (
	// BScheduleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBScheduleModel.
	BScheduleModel interface {
		bScheduleModel
	}

	customBScheduleModel struct {
		*defaultBScheduleModel
	}
)

// NewBScheduleModel returns a model for the database table.
func NewBScheduleModel(c redis.RedisConf) BScheduleModel {
	return &customBScheduleModel{
		defaultBScheduleModel: newBScheduleModel(c),
	}
}
