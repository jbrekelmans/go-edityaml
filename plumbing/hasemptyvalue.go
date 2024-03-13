package plumbing

import (
	"fmt"
	"math/big"

	goyaml "gopkg.in/yaml.v3"
)

func HasEmptyValue(node *goyaml.Node) (bool, error) {
	if node.Kind == goyaml.AliasNode {
		return false, nil
	}
	shortTag := node.ShortTag()
	switch shortTag {
	case "!!str":
		return node.Value == "", nil
	case "!!bool":
		b, err := getNodeValueAsBool(node)
		if err != nil {
			return false, err
		}
		return !b, nil
	case "!!null":
		return true, nil
	case "!!int":
		i, err := getNodeValueAsInt(node)
		if err != nil {
			return false, err
		}
		var zero big.Int
		return i.Cmp(&zero) == 0, nil
	case "!!float":
		f, err := getNodeValueAsFloat(node)
		if err != nil {
			return false, err
		}
		return f == 0.0, nil
	case "!!seq", "!!map":
		return len(node.Content) == 0, nil
	}
	return false, fmt.Errorf(`invalid node (shortTag = %#v, kind = %s)`, shortTag, nodeKind(node.Kind))
}
