package property

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// Charge type enum
const (
	Rate int64 = iota
	Fixed
)
