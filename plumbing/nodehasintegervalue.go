package plumbing

import (
	"math/big"

	goyaml "gopkg.in/yaml.v3"
)

func nodeHasIntegerValue(node *goyaml.Node, value any) (bool, error) {
	nodeValue, err := getNodeValueAsInt(node)
	if err != nil {
		return false, err
	}
	switch value := value.(type) {
	case int8:
		return nodeValue.IsInt64() && nodeValue.Int64() == int64(value), nil
	case int16:
		return nodeValue.IsInt64() && nodeValue.Int64() == int64(value), nil
	case int32:
		return nodeValue.IsInt64() && nodeValue.Int64() == int64(value), nil
	case int64:
		return nodeValue.IsInt64() && nodeValue.Int64() == int64(value), nil
	case int:
		return nodeValue.IsInt64() && nodeValue.Int64() == int64(value), nil
	case uint8:
		return nodeValue.IsUint64() && nodeValue.Uint64() == uint64(value), nil
	case uint16:
		return nodeValue.IsUint64() && nodeValue.Uint64() == uint64(value), nil
	case uint32:
		return nodeValue.IsUint64() && nodeValue.Uint64() == uint64(value), nil
	case uint64:
		return nodeValue.IsUint64() && nodeValue.Uint64() == uint64(value), nil
	case uint:
		return nodeValue.IsUint64() && nodeValue.Uint64() == uint64(value), nil
	case big.Int:
		return nodeValue.Cmp(&value) == 0, nil
	case *big.Int:
		return value != nil && nodeValue.Cmp(value) == 0, nil
	}
	return false, nil
}
