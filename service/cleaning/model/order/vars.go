package order

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// order status enum
const (
	Queuing int64 = iota
	Pending
	Working
	Unpaid
	Completed
	Cancelled
	Transfering
)
