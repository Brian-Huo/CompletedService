package payment

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BPaymentModel = (*customBPaymentModel)(nil)

type (
	// BPaymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBPaymentModel.
	BPaymentModel interface {
		bPaymentModel
		Enquire(ctx context.Context, data *BPayment) (*BPayment, error)
	}

	customBPaymentModel struct {
		*defaultBPaymentModel
	}
)

// NewBPaymentModel returns a model for the database table.
func NewBPaymentModel(conn sqlx.SqlConn, c cache.CacheConf) BPaymentModel {
	return &customBPaymentModel{
		defaultBPaymentModel: newBPaymentModel(conn, c),
	}
}

func (m *defaultBPaymentModel) Enquire(ctx context.Context, data *BPayment) (*BPayment, error) {
	pre_data, err := m.FindOneByCardNumber(ctx, data.CardNumber)
	switch err {
	case nil:
		return pre_data, nil
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
