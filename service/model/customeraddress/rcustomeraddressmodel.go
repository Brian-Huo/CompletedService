package customeraddress

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RCustomerAddressModel = (*customRCustomerAddressModel)(nil)

type (
	// RCustomerAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRCustomerAddressModel.
	RCustomerAddressModel interface {
		rCustomerAddressModel
	}

	customRCustomerAddressModel struct {
		*defaultRCustomerAddressModel
	}
)

// NewRCustomerAddressModel returns a model for the database table.
func NewRCustomerAddressModel(conn sqlx.SqlConn) RCustomerAddressModel {
	return &customRCustomerAddressModel{
		defaultRCustomerAddressModel: newRCustomerAddressModel(conn),
	}
}
