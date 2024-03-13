package plumbing

import (
	"fmt"
	"reflect"

	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/internal"
)

// NodeHasValueEqualTo determines if the node value is equal to the specified value.
// NodeHasValueEqualTo resolves alias nodes.
// NodeHasValueEqualTo only works for simple scalars (except floats) and requires strict typing:
//   - If node.ShortTag() == "!!str" and value does not have type string then returns false.
//   - If node.ShortTag() == "!!bool" and value does not have type bool then returns false.
//   - If node.ShortTag() == "!!null" and value is not nil (value != nil, and reflect.ValueOf(value).IsNil() returns false or panics) then returns false.
//   - If node.ShortTag() == "!!int" and value does not have type one of byte, (u)int8, (u)int16, (u)int32, (u)int64, (u)int, "math/big".Int and *"math/big".Int then returns false.
//   - If node.ShortTag() == "!!float" and value does not have type one of float32 and float64 then returns false.
//   - If node.ShortTag() is not one of "!!str", "!!bool", "!!null", "!!float" and "!!int" (i.e. "!!map" and "!!seq"),
//     or the node value is invalid, then returns an error.
func NodeHasValueEqualTo(node *goyaml.Node, value any) (bool, error) {
	node, err := ResolveAlias(node)
	if err != nil {
		return false, err
	}
	shortTag := node.ShortTag()
	switch shortTag {
	case "!!str":
		valueAsString, ok := value.(string)
		return ok && node.Value == valueAsString, nil
	case "!!bool":
		b, err := getNodeValueAsBool(node)
		if err != nil {
			return false, err
		}
		return b == value, nil
	case "!!null":
		if value == nil {
			return true, nil
		}
		isNil, _ := internal.IsNilSafe(reflect.ValueOf(value))
		return isNil, nil
	case "!!int":
		return nodeHasIntegerValue(node, value)
	case "!!float":
		return nodeHasFloatValue(node, value)
	}
	return false, fmt.Errorf(
		`NodeHasValueEqualTo is not implemented for node (shortTag = %#v, kind = %s, value = %#v) and value of type %T`,
		shortTag,
		nodeKind(node.Kind),
		node.Value,
		value,
	)
}
