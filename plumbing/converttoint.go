package plumbing

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

func convertBigIntToInt(v *big.Int) (int, error) {
	if !v.IsInt64() {
		return 0, errors.New(`"math/big".Int is out of range of int`)
	}
	i, ok := convertInt64ToInt(v.Int64())
	if !ok {
		return 0, errors.New(`"math/big".Int is out of range of int`)
	}
	return i, nil
}

func convertInt64ToInt(i int64) (int, bool) {
	if i >= int64(math.MinInt) && i <= int64(math.MaxInt) {
		return int(i), true
	}
	return 0, false
}

func convertUint32ToInt(i uint32) (int, error) {
	if math.MaxInt > math.MaxUint32 {
		return int(i), nil
	}
	maxInt := math.MaxInt
	if i <= uint32(maxInt) {
		return int(i), nil
	}
	return 0, errors.New("uint32 value out of range of int")
}

func convertToInt(v any) (int, error) {
	switch v := v.(type) {
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		i, ok := convertInt64ToInt(v)
		if !ok {
			return 0, errors.New("int64 value out of range of int")
		}
		return i, nil
	case int:
		return v, nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return convertUint32ToInt(v)
	case uint64:
		if v > uint64(math.MaxInt) {
			return 0, errors.New("uint64 value out of range of int")
		}
		return int(v), nil
	case uint:
		if v > uint(math.MaxInt) {
			return 0, errors.New("uint value out of range of int")
		}
		return int(v), nil
	case big.Int:
		return convertBigIntToInt(&v)
	case *big.Int:
		if v == nil {
			return 0, errors.New(`cannot convert nil *"math/big".Int to int`)
		}
		return convertBigIntToInt(v)
	}
	return 0, fmt.Errorf(`not an integer type %T`, v)
}
