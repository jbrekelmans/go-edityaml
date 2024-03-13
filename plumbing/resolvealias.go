package plumbing

import (
	"fmt"

	goyaml "gopkg.in/yaml.v3"
)

func ResolveAlias(node *goyaml.Node) (*goyaml.Node, error) {
	if node.Kind != goyaml.AliasNode {
		return node, nil
	}
	const max = 100
	i := 0
	for {
		if node.Alias == nil {
			return nil, fmt.Errorf(`invalid node: kind = %s but Alias is nil`, nodeKind(node.Kind))
		}
		node = node.Alias
		if node.Kind != goyaml.AliasNode {
			return node, nil
		}
		i += 1
		if i == max {
			// Avoid getting stuck on self-refential / circular / unreasonable structures.
			// This should never happen for valid nodes.
			return nil, fmt.Errorf(`aborting after resolving an alias node %d times`, max)
		}
	}
}
