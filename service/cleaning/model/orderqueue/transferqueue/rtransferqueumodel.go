package transferqueue

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var _ RTransferQueueModel = (*customRTransferQueueModel)(nil)

type (
	// RTransferQueueModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRTransferQueueModel.
	RTransferQueueModel interface {
		rTransferQueueModel
		List() (*map[string]string, error)
	}

	customRTransferQueueModel struct {
		*defaultRTransferQueueModel
	}
)

// NewRTransferQueueModel returns a model for the database table.
func NewRTransferQueueModel(c redis.RedisConf) RTransferQueueModel {
	return &customRTransferQueueModel{
		defaultRTransferQueueModel: newRTransferQueueModel(c),
	}
}

func (m *defaultRTransferQueueModel) List() (*map[string]string, error) {
	ret, err := m.conn.Hgetall(cacheRTransferQueuePrefix)
	switch err {
	case nil:
		return &ret, nil
	case redis.Nil:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
