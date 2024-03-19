package edityaml

import (
	"fmt"

	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/plumbing"
)

// Add appends an item to the sequence at the specified path.
func Add(node *goyaml.Node, path Path, item *goyaml.Node) (err error) {
	var i int
	if len(path) > 0 {
		node, i, err = Get(node, path)
		if err != nil {
			return
		}
		if i < len(path) {
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
				if i == len(path) {
					break
				}
			}
			node.Kind = goyaml.SequenceNode
		}
	}
	if node.Kind != goyaml.SequenceNode {
		err = fmt.Errorf(`node at path %s is not a sequence`, path)
		return
	}
	node.Content = append(node.Content, item)
	return
}
