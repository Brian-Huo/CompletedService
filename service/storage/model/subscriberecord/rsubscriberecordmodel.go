package subscriberecord

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RSubscribeRecordModel = (*customRSubscribeRecordModel)(nil)

type (
	// RSubscribeRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRSubscribeRecordModel.
	RSubscribeRecordModel interface {
		rSubscribeRecordModel
	}

	customRSubscribeRecordModel struct {
		*defaultRSubscribeRecordModel
	}
)

// NewRSubscribeRecordModel returns a model for the database table.
func NewRSubscribeRecordModel(conn sqlx.SqlConn) RSubscribeRecordModel {
	return &customRSubscribeRecordModel{
		defaultRSubscribeRecordModel: newRSubscribeRecordModel(conn),
	}
}
