package edityaml

import (
	"fmt"

	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/plumbing"
	"github.com/rs/zerolog/log"
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

func setCommon(node *goyaml.Node, path Path) (valueNode *goyaml.Node, err error) {
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
				i++
				if i == len(path)-1 {
					break
				}
			}
		}
		var j int
		if node.Kind == goyaml.MappingNode {
			log.Debug().Msgf("MAPPINGNODE FOUND")
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
				valueNode = new(goyaml.Node)
				addToMapping(node, key, valueNode)
				return
			}
		} else if node.Kind == goyaml.SequenceNode {
			log.Debug().Msgf("SEQUENCENODE FOUND")
			j, err = plumbing.AccessSequence(node, path[i])
			if err != nil {
				err = fmt.Errorf("error doing AccessSequence(<node at path %s>, %#v): %w", path[:i], path[i], err)
				return
			}
		} else {
			err = fmt.Errorf(`cannot set within node at path %s: node is neither a map nor a sequence`, path[:i])
			return
		}
		log.Debug().Msgf("VALUENODE 123 IS %#v", valueNode)
		valueNode = node.Content[j]
		log.Debug().Msgf("VALUENODE 456 IS %#v", valueNode)
	} else {
		log.Debug().Msgf("SCALARNODE FOUND")
		valueNode = node
	}
	log.Debug().Msgf("VALUENODE FINAL IS %#v", valueNode)
	return
}

// SetString sets the node at the specified path to the specified string.
// Maps are created along the path as needed.
// path can be empty, in which case node is updated to a scalar with the specified string value.
// The scalar node representing the specified string is returned.
// This can be used to control the style/comments in the output YAML.
func SetString(node *goyaml.Node, path Path, value string) (valueNode *goyaml.Node, changed bool, err error) {
	valueNode, changed, err = SetScalar(node, path, value)
	return
}

// SetBool sets the node at the specified path to the specified boolean value.
// Maps are created along the path as needed.
// path can be empty, in which case node is updated to a scalar with the specified boolean value.
// The scalar node representing the specified boolean is returned.
// This can be used to control the style/comments in the output YAML.
func SetBool(node *goyaml.Node, path Path, value bool) (valueNode *goyaml.Node, changed bool, err error) {
	valueNode, changed, err = SetScalar(node, path, value)
	return
}

// SetInt sets the node at the specified path to the specified int.
// Maps are created along the path as needed.
// path can be empty, in which case node is updated to a scalar with the specified int value.
// The scalar node representing the specified int is returned.
// This can be used to control the style/comments in the output YAML.
func SetInt(node *goyaml.Node, path Path, value int64) (valueNode *goyaml.Node, changed bool, err error) {
	valueNode, changed, err = SetScalar(node, path, value)
	return
}

// SetScalar sets a scalar value at the given path within the given node.
// The value is converted to a node via "github.com/jbrekelmans/go-edityaml/plumbing".MakeScalar.
// Maps are created along the path as needed.
// path can be empty, in which case node is updated to a scalar with the specified value.
// The scalar node representing the specified value is returned.
// This can be used to control the style/comments in the output YAML.
func SetScalar(node *goyaml.Node, path Path, value any) (valueNode *goyaml.Node, changed bool, err error) {
	valueNode, err = setCommon(node, path)
	if err != nil {
		return
	}
	ok, err := plumbing.NodeHasValueEqualTo(valueNode, value)
	if err != nil {
		return
	}
	if !ok {
		var newValueNode *goyaml.Node
		newValueNode, err = plumbing.MakeScalar(value)
		if err != nil {
			return
		}
		*valueNode = *newValueNode
		changed = true
	}
	return
}

// Set sets a node at the given path within the given node.
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
func Set(node *goyaml.Node, path Path, value goyaml.Node) (valueNode *goyaml.Node, err error) {
	valueNode, err = setCommon(node, path)
	if err != nil {
		return
	}
	*valueNode = value
	return
}
