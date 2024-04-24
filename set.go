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
		if i < len(path)-1 {
			if node.Kind != goyaml.MappingNode {
				err = fmt.Errorf(`node at path %s is not a mapping`, path[:i])
				return
			}
			for {
				// Insert new mapping nodes.
				var key *goyaml.Node
				key, err = plumbing.MakeScalar(path[i])
				if err != nil {
					err = fmt.Errorf(`error inserting into mapping: error generating key for path[%d]: path[%d] has unsupported type %T: %w`, i, i, path[i], err)
					return
				}
				value := &goyaml.Node{
					Kind: goyaml.MappingNode,
				}
				addToMapping(node, key, value)
				node = value
				changed = true
				i++
				if i == len(path)-1 {
					break
				}
			}
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

// SetBool sets the node at the specified path to the specified boolean value.
// Maps are created along the path as needed.
// path can be empty, in which case node is updated to a scalar with the specified boolean value.
// The scalar node representing the specified boolean is returned.
// This can be used to control the style/comments in the output YAML.
func SetBool(node *goyaml.Node, path Path, value bool) (valueNode *goyaml.Node, changed bool, err error) {
	valueNode, changed, err = set(node, path, value, func() *goyaml.Node {
		return plumbing.MakeBoolScalar(value)
	})
	return
}

// SetInt sets the node at the specified path to the specified int.
// Maps are created along the path as needed.
// path can be empty, in which case node is updated to a scalar with the specified int value.
// The scalar node representing the specified int is returned.
// This can be used to control the style/comments in the output YAML.
func SetInt(node *goyaml.Node, path Path, value int64) (valueNode *goyaml.Node, changed bool, err error) {
	valueNode, changed, err = set(node, path, value, func() *goyaml.Node {
		return plumbing.MakeIntScalar(value)
	})
	return
}

// SetScalar sets a scalar value at the given path within the given node.
//
// node is the node to set the value within.
//
// path is the path within the node to set the value at.
//
// value is the scalar value to set, eg. "hello".
//
// Returns a bool showing whether or not the value set updated an existing value (ie. value was changed).
func SetScalar(node *goyaml.Node, path Path, value any) (changed bool, err error) {
	n, err := plumbing.MakeScalar(value)
	if err == nil {
		_, changed, err = set(node, path, nil, func() *goyaml.Node {
			return n
		})
	}

	return
}

// SetScalar sets a node at the given path within the given node.
//
// node is the node to set the value within.
//
// path is the path within the node to set the value at.
//
// value is the node to set at the specified path.
//
// Returns the node that was added.
//
// Example, adding a new empty mapping node to "node.x.y":
//
//	value := &goyaml.Node{
//		Kind: goyaml.MappingNode,
//		Tag:  "!!map",
//	}
//
// err := Set(node, ".x.y", value)
// ...
func Set(node *goyaml.Node, path Path, value *goyaml.Node) (addedNode *goyaml.Node, err error) {
	addedNode, _, err = set(node, path, nil, func() *goyaml.Node {
		return value
	})
	return
}
