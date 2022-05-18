package contractorservice

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RContractorServiceModel = (*customRContractorServiceModel)(nil)

type (
	// RContractorServiceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRContractorServiceModel.
	RContractorServiceModel interface {
		rContractorServiceModel
	}

	customRContractorServiceModel struct {
		*defaultRContractorServiceModel
	}
)

// NewRContractorServiceModel returns a model for the database table.
func NewRContractorServiceModel(conn sqlx.SqlConn) RContractorServiceModel {
	return &customRContractorServiceModel{
		defaultRContractorServiceModel: newRContractorServiceModel(conn),
	}
}
