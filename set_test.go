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

func ExampleSetSequence() {
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
	content := []any{
		"yello",
		true,
		int64(123),
		"hello",
	}
	_, _, _ = SetSequence(node, MustParsePath(".new"), content)
	bytes, _ := goyaml.Marshal(docNode)
	fmt.Println(string(bytes))
	// Output:
	// # Comment that should be preserved.
	// key: [1, "hi"]
	// h:
	//     - abc
	//     - 123
	// new:
	//     - yello
	//     - true
	//     - 123
	//     - hello
}

func ExampleSetNewMap() {
	// Load YAML data
	docNode := new(goyaml.Node)
	_ = goyaml.Unmarshal([]byte(`---
# Comment that should be preserved.
key: [1, "hi"]
`), docNode)
	node := docNode.Content[0]

	// Edit.
	_, _, _ = SetNewMap(node, MustParsePath(".new"))
	_, _, _ = SetString(node, MustParsePath(".new[\"abc\"]"), "yello")
	bytes, _ := goyaml.Marshal(docNode)
	fmt.Println(string(bytes))
	// Output:
	// # Comment that should be preserved.
	// key: [1, "hi"]
	// new:
	//     abc: yello
}

func ExampleSetNewMap_second() {
	// Load YAML data
	docNode := new(goyaml.Node)
	_ = goyaml.Unmarshal([]byte(`---
# Comment that should be preserved.
key: [1, "hi"]
`), docNode)
	node := docNode.Content[0]

	// Edit.
	_, _, _ = SetNewMap(node, MustParsePath(".new"))
	content := []any{
		"yello",
		true,
		int64(123),
		"hello",
	}
	_, _, _ = SetString(node, MustParsePath(".new[\"abc\"]"), "yello")
	_, _, _ = SetSequence(node, MustParsePath(".new[\"seq_thing\"]"), content)
	bytes, _ := goyaml.Marshal(docNode)
	fmt.Println(string(bytes))
	// Output:
	// # Comment that should be preserved.
	// key: [1, "hi"]
	// new:
	//     abc: yello
	//     seq_thing:
	//         - yello
	//         - true
	//         - 123
	//         - hello
}
