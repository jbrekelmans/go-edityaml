package internal

import (
	"reflect"
)

// IsNilSafe returns (isNil, ok) = (rv.IsNil(), true) if rv.IsNil() does not panic.
// Returns (isNil, ok) = (false, false) otherwise.
func IsNilSafe(rv reflect.Value) (isNil, ok bool) {
	defer func() {
		recover()
	}()
	isNil = rv.IsNil()
	ok = true
	return
}
