package plumbing

import (
	"errors"
	"fmt"

	goyaml "gopkg.in/yaml.v3"
)

// DoMapLookup finds a key in a mapping node.
// If node is invalid then returns an error.
// If the mapping does not contain key then returns -1, nil.
// Key equality is defined by NodeHasValueEqualTo.
func DoMapLookup(node *goyaml.Node, key any) (int, error) {
	if node.Kind != goyaml.MappingNode {
		return -1, fmt.Errorf("node is not a mapping (kind = %s)", nodeKind(node.Kind))
	}
	if len(node.Content)%2 != 0 {
		return -1, errors.New("invalid mapping node: odd number of children")
	}
	for i := 0; i < len(node.Content); i += 2 {
		keyNode := node.Content[i]
		if keyNode == nil {
			return -1, fmt.Errorf("invalid mapping node: Content[%d] is nil", i)
		}
		equal, err := NodeHasValueEqualTo(keyNode, key)
		if err != nil {
			return -1, err
		}
		if equal {
			if node.Content[i+1] == nil {
				return -1, fmt.Errorf("invalid mapping node: Content[%d] is nil", i+1)
			}
			return i + 1, nil
		}
	}
	return -1, nil
}
