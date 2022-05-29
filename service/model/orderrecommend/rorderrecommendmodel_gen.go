// Code generated by goctl. DO NOT EDIT!

package orderrecommend

import (
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

var (
	cacheROrderRecommendContractorIdPrefix = "cache:rOrderRecommend:contractId:"
)

type (
	rOrderRecommendModel interface {
		Insert(data *ROrderRecommend) (int, error)
		FindOne(contractorId int64, orderId int64) (int, error)
		List(contractorId int64) (*[]int64, error)
		Delete(contractorId int64, orderId int64) (int, error)
		DeleteAll(contractorId int64) (int, error)
	}

	defaultROrderRecommendModel struct {
		conn redis.Redis
		table string
	}

	ROrderRecommend struct {
		ContractorId int64 `db:"contractor_id"`
		OrderId      int64 `db:"order_id"`
	}
)

func newROrderRecommendModel(c redis.RedisConf) *defaultROrderRecommendModel {
	return &defaultROrderRecommendModel{
		conn:       *c.NewRedis(),
		table:      "`r_order_recommend`",
	}
}

func (m *defaultROrderRecommendModel) Insert(data *ROrderRecommend) (int, error) {
	rOrderRecommendContractorIdKey := fmt.Sprintf("%s%v", cacheROrderRecommendContractorIdPrefix, data.ContractorId)
	ret, err := m.conn.Sadd(rOrderRecommendContractorIdKey, data.OrderId)
	return ret, err
}

func (m *defaultROrderRecommendModel) FindOne(contractorId int64, orderId int64) (int, error) {
	rOrderRecommendContractorIdKey := fmt.Sprintf("%s%v", cacheROrderRecommendContractorIdPrefix, contractorId)
	ret, _, err := m.conn.Sscan(rOrderRecommendContractorIdKey, 0, strconv.FormatInt(orderId, 10), 1)
	
	if err != nil {
		return 0, err
	} else if len(ret) == 0 {
		return 0, ErrNotFound
	}
	return 1, nil
}

func (m *defaultROrderRecommendModel) List(contractorId int64) (*[]int64, error) {
	rOrderRecommendContractorIdKey := fmt.Sprintf("%s%v", cacheROrderRecommendContractorIdPrefix, contractorId)
	ret, err := m.conn.Smembers(rOrderRecommendContractorIdKey)

	switch err {
	case nil:
		var resp []int64
		for _, val := range ret{
			order_id, err := strconv.ParseInt(val, 10, 64)
			if err == nil {
				resp = append(resp, order_id)
			}
		}
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *defaultROrderRecommendModel) Delete(contractorId int64, orderId int64) (int, error) {
	rOrderRecommendContractorIdKey := fmt.Sprintf("%s%v", cacheROrderRecommendContractorIdPrefix, contractorId)
	ret, err := m.conn.Srem(rOrderRecommendContractorIdKey, orderId)
	return ret, err
}

func (m *defaultROrderRecommendModel) DeleteAll(contractorId int64) (int, error) {
	rOrderRecommendContractorIdKey := fmt.Sprintf("%s%v", cacheROrderRecommendContractorIdPrefix, contractorId)
	ret, err := m.conn.Del(rOrderRecommendContractorIdKey)
	return ret, err
}

func (m *defaultROrderRecommendModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheROrderRecommendContractorIdPrefix, primary)
}

func (m *defaultROrderRecommendModel) tableName() string {
	return m.table
}
