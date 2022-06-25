package contractor

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// contractor workstatus enum
const (
	Vacant int64 = iota
	InWork
	Await
	InRest
	Resigned
)

// contractor type enum
const (
	Employee int64 = iota
	Individual
)
