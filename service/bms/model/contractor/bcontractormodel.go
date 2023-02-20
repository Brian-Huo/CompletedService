package contractor

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BContractorModel = (*customBContractorModel)(nil)

type (
	// BContractorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBContractorModel.
	BContractorModel interface {
		bContractorModel
		FindAllByFinance(ctx context.Context, financeId int64) ([]*BContractor, error)
		Resign(ctx context.Context, contractorId int64) error
		ResignByFinance(ctx context.Context, financeId int64) error
	}

	customBContractorModel struct {
		*defaultBContractorModel
	}
)

// NewBContractorModel returns a model for the database table.
func NewBContractorModel(conn sqlx.SqlConn, c cache.CacheConf) BContractorModel {
	return &customBContractorModel{
		defaultBContractorModel: newBContractorModel(conn, c),
	}
}

func (m *defaultBContractorModel) FindAllByFinance(ctx context.Context, financeId int64) ([]*BContractor, error) {
	var resp []*BContractor
	query := fmt.Sprintf("select %s from %s where `finance_id` = ?", bContractorRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, financeId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBContractorModel) Resign(ctx context.Context, contractorId int64) error {
	data, err := m.FindOne(ctx, contractorId)
	if err != nil {
		return err
	}

	data.WorkStatus = Resigned
	bContractorContactDetailsKey := fmt.Sprintf("%s%v", cacheBContractorContactDetailsPrefix, data.ContactDetails)
	bContractorContractorIdKey := fmt.Sprintf("%s%v", cacheBContractorContractorIdPrefix, data.ContractorId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `contractor_id` = ?", m.table, bContractorRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ContractorPhoto, data.ContractorName, data.ContractorType, data.ContactDetails, data.FinanceId, data.AddressId, data.LinkCode, data.WorkStatus, data.ContractorId)
	}, bContractorContractorIdKey, bContractorContactDetailsKey)
	return err
}

func (m *defaultBContractorModel) ResignByFinance(ctx context.Context, financeId int64) error {
	query := fmt.Sprintf("update %s set %s where `finance_id` = ?", m.table, "work_status=?")
	_, err := m.ExecNoCacheCtx(ctx, query, Resigned, financeId)
	return err
}
