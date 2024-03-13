package internal

import (
	"math/big"
	"testing"

	goyaml "gopkg.in/yaml.v3"
)

func TestParseYAML(t *testing.T, s string) *goyaml.Node {
	node := new(goyaml.Node)
	err := goyaml.Unmarshal([]byte(s), node)
	if err != nil {
		t.Fatalf("error parsing YAML test data: %v", err)
	}
	if node.Kind == goyaml.DocumentNode {
		node = node.Content[0]
	}
	return node
}

func TestParseBigInt(t *testing.T, s string) *big.Int {
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		t.Fatalf(`error parsing test data as "math/big".Int: %#v`, s)
	}
	return i
}
