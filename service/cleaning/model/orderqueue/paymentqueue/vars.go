package paymentqueue

import (
	"errors"
)

var ErrNotFound = errors.New("redis not found")
