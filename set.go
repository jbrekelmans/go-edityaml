package edityaml

import (
	"fmt"

	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/plumbing"
)

func addToMapping(node *goyaml.Node, key, value *goyaml.Node) {
	// Add to mapping.
	// Do a best effort to keep the mapping keys sorted.
	insertIndex := getMappingInsertIndex(node, key)
	// Ensure there is capacity for two new nodes.
	node.Content = append(node.Content, nil, nil)
	// Move things over by two.
	copy(node.Content[insertIndex+2:], node.Content[insertIndex:])
	// Write.
	node.Content[insertIndex] = key
	node.Content[insertIndex+1] = value
}

func getMappingInsertIndex(node, key *goyaml.Node) int {
	i := 0
	for i < len(node.Content) {
		key2 := node.Content[i]
		if key2.Kind != goyaml.ScalarNode || key.Value < key2.Value {
			break
		}
		i += 2
	}
	return i
}

func set(node *goyaml.Node, path Path, value any, valueNodeFactory func() *goyaml.Node) (valueNode *goyaml.Node, changed bool, err error) {
	var i int
	if len(path) > 0 {
		node, i, err = Get(node, path[:len(path)-1])
		if err != nil {
			return
		}
		for i < len(path)-1 {
			// Insert new mapping nodes.
			var key *goyaml.Node
			key, err = plumbing.MakeScalar(path[i])
			if err != nil {
				err = fmt.Errorf(`TODO: %w`, err)
				return
			}
			value := &goyaml.Node{
				Kind: goyaml.MappingNode,
			}
			addToMapping(node, key, value)
			node = value
			changed = true
			i++
		}
		var j int
		if node.Kind == goyaml.MappingNode {
			j, err = plumbing.DoMapLookup(node, path[i])
			if err != nil {
				err = fmt.Errorf("error doing DoMapLookup(<node at path %s>, %#v): %w", path[:i], path[i], err)
				return
			}
			if j < 0 {
				var key *goyaml.Node
				key, err = plumbing.MakeScalar(path[i])
				if err != nil {
					err = fmt.Errorf(`TODO: %w`, err)
					return
				}
				valueNode = valueNodeFactory()
				addToMapping(node, key, valueNode)
				changed = true
				return
			}
		} else if node.Kind == goyaml.SequenceNode {
			j, err = plumbing.AccessSequence(node, path[i])
			if err != nil {
				err = fmt.Errorf("error doing AccessSequence(<node at path %s>, %#v): %w", path[:i], path[i], err)
				return
			}
		} else {
			err = fmt.Errorf(`cannot set within node at path %s: node is neither a map nor a sequence`, path[:i])
			return
		}
		valueNode = node.Content[j]
	} else {
		valueNode = node
	}
	if ok, _ := plumbing.NodeHasValueEqualTo(valueNode, value); !ok {
		*valueNode = *valueNodeFactory()
		changed = true
	}
	return
}

// SetString sets the node at the specified path to the specified string.
// Maps are created along the path as needed.
// path can be empty, in which case node is updated to a scalar with the specified string value.
// The scalar node representing the specified string is returned.
// This can be used to control the style/comments in the output YAML.
func SetString(node *goyaml.Node, path Path, value string) (valueNode *goyaml.Node, changed bool, err error) {
	valueNode, changed, err = set(node, path, value, func() *goyaml.Node {
		return plumbing.MakeStringScalar(value)
	})
	return
}
