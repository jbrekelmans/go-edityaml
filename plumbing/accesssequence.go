package plumbing

import (
	"fmt"

	goyaml "gopkg.in/yaml.v3"
)

// AccessSequence gets an item of a sequence with the specified index.
func AccessSequence(node *goyaml.Node, index any) (int, error) {
	if node.Kind != goyaml.SequenceNode {
		return 0, fmt.Errorf("node is not a sequence (kind = %s)", nodeKind(node.Kind))
	}
	indexInt, err := convertToInt(index)
	if err != nil {
		return 0, err
	}
	if uint(indexInt) >= uint(len(node.Content)) {
		return 0, fmt.Errorf(`index out of bounds`)
	}
	if node.Content[indexInt] == nil {
		return 0, fmt.Errorf(`invalid sequence node: Content[%d] is nil`, indexInt)
	}
	return indexInt, nil
}
