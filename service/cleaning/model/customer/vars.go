package customer

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// customer type enum
const (
	Individual int64 = iota
	Agency
)
