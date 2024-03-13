package edityaml

import (
	"fmt"

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
