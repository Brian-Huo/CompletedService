package operation

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// operation enum
const (
	Accept int64 = iota
	Decline
	Transfer
)
