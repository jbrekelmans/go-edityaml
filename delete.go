package edityaml

import (
	"errors"
	"fmt"

	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/plumbing"
)

// Delete deletes the value at specified path.
// path must not be empty.
func Delete(node *goyaml.Node, path Path) (changed bool, err error) {
	if len(path) == 0 {
		return false, errors.New("path is empty")
	}
	var i int
	node, i, err = Get(node, path[:len(path)-1])
	if err != nil || i < len(path)-1 {
		return false, err
	}
	pathItem := path[i]
	if node.Kind == goyaml.MappingNode {
		j, err := plumbing.DoMapLookup(node, pathItem)
		if err != nil {
			return false, fmt.Errorf("error doing DoMapLookup(<node at path %s>, %#v): %w", path[:i], pathItem, err)
		}
		if j < 0 {
			return false, nil
		}
		copy(node.Content[j-1:], node.Content[j+1:])
		n := len(node.Content)
		node.Content[n-2] = nil
		node.Content[n-1] = nil
		node.Content = node.Content[:n-2]
		return true, nil
	}
	if node.Kind == goyaml.SequenceNode {
		j, err := plumbing.AccessSequence(node, pathItem)
		if err != nil {
			return false, fmt.Errorf("error doing AccessSequence(<node at path %s>, %#v): %w", path[:i], pathItem, err)
		}
		copy(node.Content[j:], node.Content[j+1:])
		n := len(node.Content)
		node.Content[n-1] = nil
		node.Content = node.Content[:n-1]
		return true, nil
	}
	return false, fmt.Errorf(`cannot Delete within node at path %s: node is neither a map nor a sequence`, path[:i])
}
