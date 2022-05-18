// Code generated by goctl. DO NOT EDIT!

package contractorservice

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	rContractorServiceFieldNames          = builder.RawFieldNames(&RContractorService{})
	rContractorServiceRows                = strings.Join(rContractorServiceFieldNames, ",")
	rContractorServiceRowsExpectAutoSet   = strings.Join(stringx.Remove(rContractorServiceFieldNames, "`create_time`", "`update_time`"), ",")
	rContractorServiceRowsWithPlaceHolder = strings.Join(stringx.Remove(rContractorServiceFieldNames, "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	rContractorServiceModel interface {
		Insert(ctx context.Context, data *RContractorService) (sql.Result, error)
		FindOne(ctx context.Context, contractorId int64, serviceId int64) (*RContractorService, error)
		FindAllByContractor(ctx context.Context, contractorId int64) ([]*RContractorService, error)
		FindAllByService(ctx context.Context, serviceId int64) ([]*RContractorService, error)
		Delete(ctx context.Context, contractorId int64, serviceId int64) error
		DeleteAllByContractor(ctx context.Context, contractorId int64) error
		DeleteAllByService(ctx context.Context, contractorId int64) error
	}

	defaultRContractorServiceModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RContractorService struct {
		ContractorId int64 `db:"contractor_id"`
		ServiceId    int64 `db:"service_id"`
	}
)

func newRContractorServiceModel(conn sqlx.SqlConn) *defaultRContractorServiceModel {
	return &defaultRContractorServiceModel{
		conn:  conn,
		table: "`r_contractor_service`",
	}
}

func (m *defaultRContractorServiceModel) Insert(ctx context.Context, data *RContractorService) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, rContractorServiceRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ContractorId, data.ServiceId)
	return ret, err
}

func (m *defaultRContractorServiceModel) FindOne(ctx context.Context, contractorId int64, serviceId int64) (*RContractorService, error) {
	query := fmt.Sprintf("select %s from %s where `contractor_id` = ? and `service_id` = ? limit 1", rContractorServiceRows, m.table)
	var resp RContractorService
	err := m.conn.QueryRowCtx(ctx, &resp, query, contractorId, serviceId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRContractorServiceModel) FindAllByContractor(ctx context.Context, contractorId int64) ([]*RContractorService, error) {
	query := fmt.Sprintf("select %s from %s where `contractor_id` = ?", rContractorServiceRows, m.table)
	var resp []*RContractorService
	err := m.conn.QueryRowsCtx(ctx, &resp, query, contractorId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRContractorServiceModel) FindAllByService(ctx context.Context, serviceId int64) ([]*RContractorService, error) {
	query := fmt.Sprintf("select %s from %s where `service_id` = ?", rContractorServiceRows, m.table)
	var resp []*RContractorService
	err := m.conn.QueryRowsCtx(ctx, &resp, query, serviceId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRContractorServiceModel) Delete(ctx context.Context, contractorId int64, serviceId int64) error {
	query := fmt.Sprintf("delete from %s where `contractor_id` = ? and `service_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, contractorId, serviceId)
	return err
}

func (m *defaultRContractorServiceModel) DeleteAllByContractor(ctx context.Context, contractorId int64) error {
	query := fmt.Sprintf("delete from %s where `contractor_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, contractorId)
	return err
}

func (m *defaultRContractorServiceModel) DeleteAllByService(ctx context.Context, serviceId int64) error {
	query := fmt.Sprintf("delete from %s where `service_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, serviceId)
	return err
}

func (m *defaultRContractorServiceModel) tableName() string {
	return m.table
}
