package plumbing

import (
	"fmt"
	"math/big"
	"regexp"

	goyaml "gopkg.in/yaml.v3"
)

var regexpYAMLInt = regexp.MustCompile(`^0$|^[+-]?[1-9][0-9]*$`)

func getNodeValueAsInt(node *goyaml.Node) (*big.Int, error) {
	if !regexpYAMLInt.MatchString(node.Value) {
		return nil, fmt.Errorf(`invalid node (shortTag = "!!int", value = %#v)`, node.Value)
	}
	// Given the regexp check, this should always succeed.
	i, _ := new(big.Int).SetString(node.Value, 10)
	return i, nil
}
