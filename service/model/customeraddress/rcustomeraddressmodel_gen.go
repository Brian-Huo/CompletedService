// Code generated by goctl. DO NOT EDIT!

package customeraddress

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	rCustomerAddressFieldNames          = builder.RawFieldNames(&RCustomerAddress{})
	rCustomerAddressRows                = strings.Join(rCustomerAddressFieldNames, ",")
	rCustomerAddressRowsExpectAutoSet   = strings.Join(stringx.Remove(rCustomerAddressFieldNames, "`create_time`", "`update_time`"), ",")
)

type (
	rCustomerAddressModel interface {
		Insert(ctx context.Context, data *RCustomerAddress) (sql.Result, error)
		FindOne(ctx context.Context, customerId int64, addressId int64) (*RCustomerAddress, error)
		FindAllByCustomer(ctx context.Context, customerId int64) ([]*RCustomerAddress, error)
		Delete(ctx context.Context, customerId int64, addressId int64) error
		DeleteByCustomer(ctx context.Context, customerId int64) error
	}

	defaultRCustomerAddressModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RCustomerAddress struct {
		CustomerId int64        `db:"customer_id"`
		AddressId  int64        `db:"address_id"`
		UpdateDate time.Time 	`db:"update_date"`
	}
)

func newRCustomerAddressModel(conn sqlx.SqlConn) *defaultRCustomerAddressModel {
	return &defaultRCustomerAddressModel{
		conn:  conn,
		table: "`r_customer_address`",
	}
}

func (m *defaultRCustomerAddressModel) Insert(ctx context.Context, data *RCustomerAddress) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, rCustomerAddressRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CustomerId, data.AddressId, data.UpdateDate)
	return ret, err
}

func (m *defaultRCustomerAddressModel) FindOne(ctx context.Context, customerId int64, addressId int64) (*RCustomerAddress, error) {
	query := fmt.Sprintf("select %s from %s where `customer_id` = ? and `address_id` = ? limit 1", rCustomerAddressRows, m.table)
	var resp RCustomerAddress
	err := m.conn.QueryRowCtx(ctx, &resp, query, customerId, addressId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRCustomerAddressModel) FindAllByCustomer(ctx context.Context, customerId int64) ([]*RCustomerAddress, error) {
	query := fmt.Sprintf("select %s from %s where `customer_id` = ?", rCustomerAddressRows, m.table)
	var resp []*RCustomerAddress
	err := m.conn.QueryRowsCtx(ctx, &resp, query, customerId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRCustomerAddressModel) Delete(ctx context.Context, customerId int64, addressId int64) error {
	query := fmt.Sprintf("delete from %s where `customer_id` = ? and `address_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, customerId, addressId)
	return err
}

func (m *defaultRCustomerAddressModel) DeleteByCustomer(ctx context.Context, customerId int64) error {
	query := fmt.Sprintf("delete from %s where `customer_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, customerId)
	return err
}

func (m *defaultRCustomerAddressModel) tableName() string {
	return m.table
}
