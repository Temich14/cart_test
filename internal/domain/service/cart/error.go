package cart

import "errors"

var (
	ErrQuantityLessThanZero = errors.New("quantity must be greater than zero")
)
