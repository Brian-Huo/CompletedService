package customer

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BCustomerModel = (*customBCustomerModel)(nil)

type (
	// BCustomerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBCustomerModel.
	BCustomerModel interface {
		bCustomerModel
		Enquire(ctx context.Context, data *BCustomer) (*BCustomer, error)
	}

	customBCustomerModel struct {
		*defaultBCustomerModel
	}
)

// NewBCustomerModel returns a model for the database table.
func NewBCustomerModel(conn sqlx.SqlConn, c cache.CacheConf) BCustomerModel {
	return &customBCustomerModel{
		defaultBCustomerModel: newBCustomerModel(conn, c),
	}
}

func (m *defaultBCustomerModel) Enquire(ctx context.Context, data *BCustomer) (*BCustomer, error) {
	pre_data, err := m.FindOneByCustomerPhone(ctx, data.CustomerPhone)
	switch err {
	case nil:
		data.CustomerId = pre_data.CustomerId
		data.CustomerType = pre_data.CustomerType
		err = m.Update(ctx, data)
		if err != nil {
			return nil, err
		}
		return data, nil
	case sqlc.ErrNotFound:
		ret, err := m.Insert(ctx, data)
		if err != nil {
			return nil, err
		}
		enquiryId, err := ret.LastInsertId()
		if err != nil {
			return nil, err
		}
		return m.FindOne(ctx, enquiryId)
	default:
		return nil, err
	}
}
