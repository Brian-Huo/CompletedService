package region

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BRegionModel = (*customBRegionModel)(nil)

type (
	// BRegionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBRegionModel.
	BRegionModel interface {
		bRegionModel
		Enquire(ctx context.Context, data *BRegion) (*BRegion, error)
	}

	customBRegionModel struct {
		*defaultBRegionModel
	}
)

// NewBRegionModel returns a model for the database table.
func NewBRegionModel(conn sqlx.SqlConn, c cache.CacheConf) BRegionModel {
	return &customBRegionModel{
		defaultBRegionModel: newBRegionModel(conn, c),
	}
}

func (m *defaultBRegionModel) Enquire(ctx context.Context, data *BRegion) (*BRegion, error) {
	pre_data, err := m.FindOneByPostcode(ctx, data.Postcode)
	switch err {
	case nil:
		return pre_data, nil
	case ErrNotFound:
		ret, err := m.Insert(ctx, data)
		if err != nil {
			return nil, err
		}
		data.RegionId, err = ret.LastInsertId()
		if err != nil {
			return nil, err
		}
		return data, nil
	default:
		return nil, err
	}
}
