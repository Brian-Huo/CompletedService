package awaitqueue

import (
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

var _ RAwaitQueueModel = (*customRAwaitQueueModel)(nil)

type (
	// RAwaitQueueModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRAwaitQueueModel.
	RAwaitQueueModel interface {
		rAwaitQueueModel
		Count() error
		List() (*map[string]string, error)
	}

	customRAwaitQueueModel struct {
		*defaultRAwaitQueueModel
	}
)

// NewRAwaitQueueModel returns a model for the database table.
func NewRAwaitQueueModel(c redis.RedisConf) RAwaitQueueModel {
	return &customRAwaitQueueModel{
		defaultRAwaitQueueModel: newRAwaitQueueModel(c),
	}
}

func (m *defaultRAwaitQueueModel) Count() error {
	ret, err := m.conn.Hgetall(cacheRAwaitQueuePrefix)
	if err != nil {
		return err
	}

	for k, v := range ret {
		v_int, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		ret[k] = strconv.Itoa(v_int + 1)
	}

	err = m.conn.Hmset(cacheRAwaitQueuePrefix, ret)
	return err
}

func (m *defaultRAwaitQueueModel) List() (*map[string]string, error) {
	ret, err := m.conn.Hgetall(cacheRAwaitQueuePrefix)
	switch err {
	case nil:
		return &ret, nil
	case redis.Nil:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
