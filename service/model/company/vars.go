package company

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// company status enum
const (
	Abolished int64 = 0
	Active    int64 = 1
)
