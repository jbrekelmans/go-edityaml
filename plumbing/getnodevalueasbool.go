package plumbing

import (
	"fmt"

	goyaml "gopkg.in/yaml.v3"
)

func getNodeValueAsBool(node *goyaml.Node) (b bool, err error) {
	if node.Value == "true" {
		return true, nil
	}
	if node.Value == "false" {
		return false, nil
	}
	return false, fmt.Errorf(`invalid node (shortTag = "!!bool", value = %#v)`, node.Value)
}
