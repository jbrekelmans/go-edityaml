package plumbing

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"

	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/internal"
)

func MakeBigIntScalar(i *big.Int) *goyaml.Node {
	if i != nil {
		return &goyaml.Node{
			Kind:  goyaml.ScalarNode,
			Tag:   "!!int",
			Value: i.String(),
		}
	}
	return &goyaml.Node{
		Kind: goyaml.ScalarNode,
		Tag:  "!!null",
	}
}

func MakeBoolScalar(b bool) *goyaml.Node {
	return &goyaml.Node{
		Kind:  goyaml.ScalarNode,
		Tag:   "!!bool",
		Value: strconv.FormatBool(b),
	}
}

func formatFloat(f float64) string {
	if math.IsNaN(f) {
		return ".nan"
	}
	positiveInfinity := math.Inf(1)
	if f < 0 {
		if f == -positiveInfinity {
			return "-.inf"
		}
	} else if f == positiveInfinity {
		return ".inf"
	}
	// Format 'e' always aligns with the YAML spec, but is not always
	// canonical.
	// Always choose bitSize 64, even if the caller passed float32 type.
	return strconv.FormatFloat(f, 'e', -1, 64)
}

func MakeFloatScalar(f float64) *goyaml.Node {
	return &goyaml.Node{
		Kind:  goyaml.ScalarNode,
		Tag:   "!!float",
		Value: formatFloat(f),
	}
}

func MakeIntScalar(i int64) *goyaml.Node {
	return &goyaml.Node{
		Kind:  goyaml.ScalarNode,
		Tag:   "!!int",
		Value: strconv.FormatInt(i, 10),
	}
}

func MakeNullScalar() *goyaml.Node {
	return &goyaml.Node{
		Kind: goyaml.ScalarNode,
		Tag:  "!!null",
	}
}

func MakeScalar(v any) (*goyaml.Node, error) {
	switch v := v.(type) {
	case string:
		return MakeStringScalar(v), nil
	case int8:
		return MakeIntScalar(int64(v)), nil
	case int16:
		return MakeIntScalar(int64(v)), nil
	case int32:
		return MakeIntScalar(int64(v)), nil
	case int64:
		return MakeIntScalar(v), nil
	case int:
		return MakeIntScalar(int64(v)), nil
	case uint8:
		return MakeUintScalar(uint64(v)), nil
	case uint16:
		return MakeUintScalar(uint64(v)), nil
	case uint32:
		return MakeUintScalar(uint64(v)), nil
	case uint64:
		return MakeUintScalar(v), nil
	case uint:
		return MakeUintScalar(uint64(v)), nil
	case big.Int:
		return MakeBigIntScalar(&v), nil
	case *big.Int:
		return MakeBigIntScalar(v), nil
	case bool:
		return MakeBoolScalar(v), nil
	case nil:
		return MakeNullScalar(), nil
	case float64:
		return MakeFloatScalar(v), nil
	case float32:
		return MakeFloatScalar(float64(v)), nil
	default:
		if isNil, _ := internal.IsNilSafe(reflect.ValueOf(v)); isNil {
			return MakeNullScalar(), nil
		}
	}
	// TODO support float.
	return nil, fmt.Errorf(`MakeScalar on non-nil value of type %T is not supported`, v)
}

func MakeStringScalar(v string) *goyaml.Node {
	return &goyaml.Node{
		Kind:  goyaml.ScalarNode,
		Value: v,
	}
}

func MakeUintScalar(ui uint64) *goyaml.Node {
	return &goyaml.Node{
		Kind:  goyaml.ScalarNode,
		Tag:   "!!int",
		Value: strconv.FormatUint(ui, 10),
	}
}

func MustMakeScalar(v any) *goyaml.Node {
	n, err := MakeScalar(v)
	if err != nil {
		panic(err)
	}
	return n
}
