// Code generated by goctl. DO NOT EDIT!

package contractor

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	bContractorFieldNames          = builder.RawFieldNames(&BContractor{})
	bContractorRows                = strings.Join(bContractorFieldNames, ",")
	bContractorRowsExpectAutoSet   = strings.Join(stringx.Remove(bContractorFieldNames, "`contractor_id`", "`create_time`", "`update_time`"), ",")
	bContractorRowsWithPlaceHolder = strings.Join(stringx.Remove(bContractorFieldNames, "`contractor_id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBContractorContractorIdPrefix   = "cache:bContractor:contractorId:"
	cacheBContractorContactDetailsPrefix = "cache:bContractor:contactDetails:"
)

type (
	bContractorModel interface {
		Insert(ctx context.Context, data *BContractor) (sql.Result, error)
		FindOne(ctx context.Context, contractorId int64) (*BContractor, error)
		FindOneByContactDetails(ctx context.Context, contactDetails string) (*BContractor, error)
		FindAllByFinance(ctx context.Context, financeId int64) ([]*BContractor, error)
		ListVacant(ctx context.Context) ([]int64, error)
		Update(ctx context.Context, data *BContractor) error
		Resign(ctx context.Context, contractorId int64) error
		ResignByFinance(ctx context.Context, financeId int64) error
		Delete(ctx context.Context, contractorId int64) error
		DeleteAllByFinance(ctx context.Context, financeId int64) error
	}

	defaultBContractorModel struct {
		sqlc.CachedConn
		table string
	}

	BContractor struct {
		ContractorId    int64          `db:"contractor_id"`
		ContractorPhoto sql.NullString `db:"contractor_photo"`
		ContractorName  string         `db:"contractor_name"`
		ContractorType  int64          `db:"contractor_type"`
		ContactDetails  string         `db:"contact_details"`
		FinanceId       int64          `db:"finance_id"`
		AddressId       sql.NullInt64  `db:"address_id"`
		LinkCode        string         `db:"link_code"`
		WorkStatus      int64          `db:"work_status"`
		OrderId         sql.NullInt64  `db:"order_id"`
	}
)

func newBContractorModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBContractorModel {
	return &defaultBContractorModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`b_contractor`",
	}
}

func (m *defaultBContractorModel) Insert(ctx context.Context, data *BContractor) (sql.Result, error) {
	bContractorContractorIdKey := fmt.Sprintf("%s%v", cacheBContractorContractorIdPrefix, data.ContractorId)
	bContractorContactDetailsKey := fmt.Sprintf("%s%v", cacheBContractorContactDetailsPrefix, data.ContactDetails)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, bContractorRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ContractorPhoto, data.ContractorName, data.ContractorType, data.ContactDetails, data.FinanceId, data.AddressId, data.LinkCode, data.WorkStatus, data.OrderId)
	}, bContractorContractorIdKey, bContractorContactDetailsKey)
	return ret, err
}

func (m *defaultBContractorModel) FindOne(ctx context.Context, contractorId int64) (*BContractor, error) {
	bContractorContractorIdKey := fmt.Sprintf("%s%v", cacheBContractorContractorIdPrefix, contractorId)
	var resp BContractor
	err := m.QueryRowCtx(ctx, &resp, bContractorContractorIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `contractor_id` = ? limit 1", bContractorRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, contractorId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBContractorModel) FindOneByContactDetails(ctx context.Context, contactDetails string) (*BContractor, error) {
	bContractorContactDetailsKey := fmt.Sprintf("%s%v", cacheBContractorContactDetailsPrefix, contactDetails)
	var resp BContractor
	err := m.QueryRowCtx(ctx, &resp, bContractorContactDetailsKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `contact_details` = ? limit 1", bContractorRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, contactDetails)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
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

func (m *defaultBContractorModel) ListVacant(ctx context.Context) ([]int64, error) {
	var resp []int64

	query := fmt.Sprintf("select %s from %s where `work_status` = ?", "contractor_id", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, Vacant)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBContractorModel) Update(ctx context.Context, data *BContractor) error {
	bContractorContactDetailsKey := fmt.Sprintf("%s%v", cacheBContractorContactDetailsPrefix, data.ContactDetails)
	bContractorContractorIdKey := fmt.Sprintf("%s%v", cacheBContractorContractorIdPrefix, data.ContractorId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `contractor_id` = ?", m.table, bContractorRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ContractorPhoto, data.ContractorName, data.ContractorType, data.ContactDetails, data.FinanceId, data.AddressId, data.LinkCode, data.WorkStatus, data.OrderId, data.ContractorId)
	}, bContractorContractorIdKey, bContractorContactDetailsKey)
	return err
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
		return conn.ExecCtx(ctx, query, data.ContractorPhoto, data.ContractorName, data.ContractorType, data.ContactDetails, data.FinanceId, data.AddressId, data.LinkCode, data.WorkStatus, data.OrderId, data.ContractorId)
	}, bContractorContractorIdKey, bContractorContactDetailsKey)
	return err
}

func (m *defaultBContractorModel) ResignByFinance(ctx context.Context, financeId int64) error {
	query := fmt.Sprintf("update %s set %s where `finance_id` = ?", m.table, "work_status=?")
	_,  err := m.ExecNoCacheCtx(ctx, query, Resigned, financeId)
	return err
}

func (m *defaultBContractorModel) Delete(ctx context.Context, contractorId int64) error {
	data, err := m.FindOne(ctx, contractorId)
	if err != nil {
		return err
	}

	bContractorContractorIdKey := fmt.Sprintf("%s%v", cacheBContractorContractorIdPrefix, contractorId)
	bContractorContactDetailsKey := fmt.Sprintf("%s%v", cacheBContractorContactDetailsPrefix, data.ContactDetails)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `contractor_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, contractorId)
	}, bContractorContactDetailsKey, bContractorContractorIdKey)
	return err
}

func (m *defaultBContractorModel) DeleteAllByFinance(ctx context.Context, financeId int64) error {
	query := fmt.Sprintf("delete from %s where `finance_id` = ?", m.table)
	_,  err := m.ExecNoCacheCtx(ctx, query, financeId)
	return err
}

func (m *defaultBContractorModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBContractorContractorIdPrefix, primary)
}

func (m *defaultBContractorModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `contractor_id` = ? limit 1", bContractorRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBContractorModel) tableName() string {
	return m.table
}
