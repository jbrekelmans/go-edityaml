package edityaml

import (
	"fmt"

	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/plumbing"
)

// Get gets a node at the specified path.
// If len(path) == 0 then returns (value, i, err) = (node, 0, nil).
// i is the number of path items walked.
// If the value was not found then i < len(path).
// If the value was found then i == len(path).
func Get(node *goyaml.Node, path Path) (value *goyaml.Node, i int, err error) {
	node, err = plumbing.ResolveAlias(node)
	if err != nil {
		return
	}
	for i < len(path) {
		pathItem := path[i]
		if node.Kind == goyaml.MappingNode {
			var j int
			j, err = plumbing.DoMapLookup(node, pathItem)
			if err != nil {
				err = fmt.Errorf("error doing DoMapLookup(<node at path %s>, %#v): %w", path[:i], pathItem, err)
				return
			}
			if j < 0 {
				value = node
				return
			}
			node = node.Content[j]
		} else if node.Kind == goyaml.SequenceNode {
			var j int
			j, err = plumbing.AccessSequence(node, pathItem)
			if err != nil {
				err = fmt.Errorf("error doing AccessSequence(<node at path %s>, %#v): %w", path[:i], pathItem, err)
				return
			}
			node = node.Content[j]
		} else {
			err = fmt.Errorf(`cannot run DoMapLookup/AccessSequence on node that is neither a map nor a sequence at path %s`, path[:i])
			return
		}
		i++
		node, err = plumbing.ResolveAlias(node)
		if err != nil {
			err = fmt.Errorf("error doing ResolveAlias(<node at path %s>): %w", path[:i], err)
			return
		}
	}
	value = node
	return
}

// GetString gets a string at the specified path.
// i is the number of path items walked. If the value was found then i == len(path).
// If the value was not found then i < len(path).
// Returns an error if the node at the specified path is not a string.
func GetString(node *goyaml.Node, path Path) (s string, i int, err error) {
	value, i, err := Get(node, path)
	if err != nil || i < len(path) {
		return
	}
	shortTag := value.ShortTag()
	if shortTag != "!!str" {
		err = fmt.Errorf(`node at path %s does not have short tag "!!str"`, path)
		return
	}
	s = node.Value
	return
}
