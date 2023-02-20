package paymentqueue

import (
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

var _ RPaymentQueueModel = (*customRPaymentQueueModel)(nil)

type (
	// RPaymentQueueModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRPaymentQueueModel.
	RPaymentQueueModel interface {
		rPaymentQueueModel
		List() (*map[string]string, error)
	}

	customRPaymentQueueModel struct {
		*defaultRPaymentQueueModel
	}
)

// NewRPaymentQueueModel returns a model for the database table.
func NewRPaymentQueueModel(c redis.RedisConf) RPaymentQueueModel {
	return &customRPaymentQueueModel{
		defaultRPaymentQueueModel: newRPaymentQueueModel(c),
	}
}

func (m *defaultRPaymentQueueModel) Count() error {
	ret, err := m.conn.Hgetall(cacheRPaymentQueuePrefix)
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

	err = m.conn.Hmset(cacheRPaymentQueuePrefix, ret)
	return err
}

func (m *defaultRPaymentQueueModel) List() (*map[string]string, error) {
	ret, err := m.conn.Hgetall(cacheRPaymentQueuePrefix)
	switch err {
	case nil:
		return &ret, nil
	case redis.Nil:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
