package employeeservice

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ REmployeeServiceModel = (*customREmployeeServiceModel)(nil)

type (
	// REmployeeServiceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customREmployeeServiceModel.
	REmployeeServiceModel interface {
		rEmployeeServiceModel
	}

	customREmployeeServiceModel struct {
		*defaultREmployeeServiceModel
	}
)

// NewREmployeeServiceModel returns a model for the database table.
func NewREmployeeServiceModel(conn sqlx.SqlConn) REmployeeServiceModel {
	return &customREmployeeServiceModel{
		defaultREmployeeServiceModel: newREmployeeServiceModel(conn),
	}
}
