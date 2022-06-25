package operation

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// operation enum
const (
	Decline int64 = iota
	Accept
	Transfer
)
