package address

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BAddressModel = (*customBAddressModel)(nil)

type (
	// BAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBAddressModel.
	BAddressModel interface {
		bAddressModel
		Enquire(ctx context.Context, data *BAddress) error
	}

	customBAddressModel struct {
		*defaultBAddressModel
	}
)

// NewBAddressModel returns a model for the database table.
func NewBAddressModel(conn sqlx.SqlConn, c cache.CacheConf) BAddressModel {
	return &customBAddressModel{
		defaultBAddressModel: newBAddressModel(conn, c),
	}
}

func (m *defaultBAddressModel) Enquire(ctx context.Context, data *BAddress) error {
	err := ErrNotFound
	switch err {
	case nil:
		// data.AddressId = pre_data.AddressId
		return nil
	case ErrNotFound:
		ret, err := m.Insert(ctx, data)
		if err != nil {
			return err
		}
		data.AddressId, err = ret.LastInsertId()
		if err != nil {
			return err
		}
		return nil
	default:
		return err
	}
}
