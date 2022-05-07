package customerpayment

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RCustomerPaymentModel = (*customRCustomerPaymentModel)(nil)

type (
	// RCustomerPaymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRCustomerPaymentModel.
	RCustomerPaymentModel interface {
		rCustomerPaymentModel
	}

	customRCustomerPaymentModel struct {
		*defaultRCustomerPaymentModel
	}
)

// NewRCustomerPaymentModel returns a model for the database table.
func NewRCustomerPaymentModel(conn sqlx.SqlConn) RCustomerPaymentModel {
	return &customRCustomerPaymentModel{
		defaultRCustomerPaymentModel: newRCustomerPaymentModel(conn),
	}
}
