package edityaml

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	goyaml "gopkg.in/yaml.v3"
)

func getYAMLNode(t *testing.T, yaml string) *goyaml.Node {
	node := new(goyaml.Node)
	decoder := goyaml.NewDecoder(strings.NewReader(yaml))
	if err := decoder.Decode(node); err != nil {
		t.Fatalf("invalid YAML test data: %v", err)
	}
	var extraDoc goyaml.Node
	if err := decoder.Decode(&extraDoc); err != nil {
		if !errors.Is(err, io.EOF) {
			t.Fatalf("invalid YAML test data: unexpected error decoding second document: %v", err)
		}
	} else {
		t.Fatalf("invalid YAML test data: multiple documents")
	}
	if node.Kind != goyaml.DocumentNode {
		t.Fatalf("decoding YAML unexpectedly did not produce a document node")
	}
	return node.Content[0]
}

func Test_GetString(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		node := getYAMLNode(t, `---
key: value
`)
		s, i, err := GetString(node, Path{"key"})
		if assert.NoError(t, err) {
			assert.Equalf(t, 1, i, "i")
			assert.Equalf(t, "value", s, "s")
		}
	})
	t.Run("NotFound", func(t *testing.T) {
		node := getYAMLNode(t, `---
key: value
`)
		_, i, err := GetString(node, Path{"key2"})
		if assert.NoError(t, err) {
			assert.Equalf(t, 0, i, "i")
		}
	})
	t.Run("TypeError1", func(t *testing.T) {
		node := getYAMLNode(t, `---
- item1
- item2
`)
		_, _, err := GetString(node, Path{"key"})
		assert.ErrorContains(t, err, "error doing AccessSequence")
	})
	t.Run("TypeError2", func(t *testing.T) {
		node := getYAMLNode(t, `---
key: 1
`)
		_, _, err := GetString(node, Path{"key"})
		assert.ErrorContains(t, err, `node at path .key does not have short tag "!!str"`)
	})
}
