package edityaml

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	goyaml "gopkg.in/yaml.v3"
)

func ExampleSetString() {
	// Load YAML data
	docNode := new(goyaml.Node)
	_ = goyaml.Unmarshal([]byte(`---
# Comment that should be preserved.
key: [1, "hi"]
`), docNode)
	node := docNode.Content[0]

	// Edit.
	_, changed, _ := SetString(node, MustParsePath(".key[1]"), "hello")
	fmt.Println(changed)
	bytes, _ := goyaml.Marshal(docNode)
	fmt.Println(string(bytes))
	// Output:
	// true
	// # Comment that should be preserved.
	// key: [1, hello]
}

func ExampleSetString_second() {
	// Load YAML data
	docNode := new(goyaml.Node)
	_ = goyaml.Unmarshal([]byte(`---
# Comment that should be preserved.
key: [1, "hi"]
`), docNode)
	node := docNode.Content[0]

	// Edit.
	_, _, _ = SetString(node, MustParsePath(".new"), "1")
	bytes, _ := goyaml.Marshal(docNode)
	fmt.Println(string(bytes))
	// Output:
	// # Comment that should be preserved.
	// key: [1, "hi"]
	// new: 1
}

func Test_SetScalar(t *testing.T) {
	t.Run("WorksIfValueDoesNotExistYet", func(t *testing.T) {
		node := getYAMLNode(t, `---
key1: {}
`)
		valueNode, changed, err := SetScalar(node, Path{"key1", "key2"}, true)
		if assert.NoError(t, err) {
			assert.Equalf(t, true, changed, "changed is true")
			assert.Equalf(t, goyaml.ScalarNode, valueNode.Kind, `valueNode has Kind scalar`)
			assert.Equalf(t, "!!bool", valueNode.ShortTag(), `valueNode has short tag "!!bool"`)
			assert.Equalf(t, "true", valueNode.Value, `valueNode.Value equals "true"`)
		}
	})
	t.Run("WorksIfValueAlreadyExists1", func(t *testing.T) {
		node := getYAMLNode(t, `---
key: asdf
`)
		valueNode, changed, err := SetScalar(node, Path{"key"}, true)
		if assert.NoError(t, err) {
			assert.Equalf(t, true, changed, "changed is true")
			assert.Equalf(t, goyaml.ScalarNode, valueNode.Kind, `valueNode has Kind scalar`)
			assert.Equalf(t, "!!bool", valueNode.ShortTag(), `valueNode has short tag "!!bool"`)
			assert.Equalf(t, "true", valueNode.Value, `valueNode.Value equals "true"`)
		}
	})
	t.Run("WorksIfValueAlreadyExists2", func(t *testing.T) {
		node := getYAMLNode(t, `---
key: true
`)
		valueNode, changed, err := SetScalar(node, Path{"key"}, true)
		if assert.NoError(t, err) {
			assert.Equalf(t, false, changed, "changed is false")
			assert.Equalf(t, goyaml.ScalarNode, valueNode.Kind, `valueNode has Kind scalar`)
			assert.Equalf(t, "!!bool", valueNode.ShortTag(), `valueNode has short tag "!!bool"`)
			assert.Equalf(t, "true", valueNode.Value, `valueNode.Value equals "false"`)
		}
	})
}

func ExampleSetScalar() {
	// Load YAML data
	docNode := new(goyaml.Node)
	_ = goyaml.Unmarshal([]byte(`---
# Comment that should be preserved.
key: [1, "hi"]
h:
- abc
- 123
`), docNode)
	node := docNode.Content[0]

	// Edit.
	_, _, _ = SetScalar(node, MustParsePath(".new"), "some content")
	bytes, _ := goyaml.Marshal(docNode)
	fmt.Println(string(bytes))
	// Output:
	// # Comment that should be preserved.
	// key: [1, "hi"]
	// h:
	//     - abc
	//     - 123
	// new: some content
}

func ExampleSet() {
	// Load YAML data
	docNode := new(goyaml.Node)
	_ = goyaml.Unmarshal([]byte(`---
# Comment that should be preserved.
key: [1, "hi"]
h:
- abc
- 123
`), docNode)
	node := docNode.Content[0]

	// Edit.
	content := goyaml.Node{
		Kind: goyaml.SequenceNode,
		Tag:  "!!seq",
	}
	addedNode, _ := Set(node, MustParsePath(".new"), content) // Example of setting a sequence node and adding to its contents
	addedNode.Content = []*goyaml.Node{
		{
			Kind:  goyaml.ScalarNode,
			Value: "hello there",
			Tag:   "!!str",
		},
	}
	addedNode.Content = append(addedNode.Content, &goyaml.Node{
		Kind:  goyaml.ScalarNode,
		Value: "1234",
		Tag:   "!!int",
	})
	bytes, _ := goyaml.Marshal(docNode)
	fmt.Println(string(bytes))
	// Output:
	// # Comment that should be preserved.
	// key: [1, "hi"]
	// h:
	//     - abc
	//     - 123
	// new:
	//     - hello there
	//     - 1234
}

func ExampleSet_second() {
	// Load YAML data
	docNode := new(goyaml.Node)
	_ = goyaml.Unmarshal([]byte(`---
# Comment that should be preserved.
key: [1, "hi"]
h:
- abc
- 123
`), docNode)
	node := docNode.Content[0]

	// Edit.
	content := &goyaml.Node{
		Kind: goyaml.SequenceNode,
		Tag:  "!!seq",
	}
	addedNode, _ := Set(node, MustParsePath(".new"), content) // Example of setting a sequence node and adding to its contents
	addedNode.Content = []*goyaml.Node{
		{
			Kind:  goyaml.ScalarNode,
			Value: "hello there",
			Tag:   "!!str",
		},
	}
	addedNode.Content = append(addedNode.Content, &goyaml.Node{
		Kind:  goyaml.ScalarNode,
		Value: "1234",
		Tag:   "!!int",
	})
	bytes, _ := goyaml.Marshal(docNode)
	fmt.Println(string(bytes))
	// Output:
	// # Comment that should be preserved.
	// key: [1, "hi"]
	// h:
	//     - abc
	//     - 123
	// new:
	//     - hello there
	//     - 1234
}
