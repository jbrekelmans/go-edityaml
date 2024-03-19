package edityaml

import (
	"fmt"

	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/plumbing"
)

func ExampleAdd() {
	// Load YAML data
	docNode := new(goyaml.Node)
	_ = goyaml.Unmarshal([]byte(`---
# Comment that should be preserved.
key: [1, "hi"]
`), docNode)
	node := docNode.Content[0]

	// Edit.
	_ = Add(node, MustParsePath(".key"), plumbing.MustMakeScalar("asdf"))
	bytes, _ := goyaml.Marshal(docNode)
	fmt.Println(string(bytes))
	// Output:
	// # Comment that should be preserved.
	// key: [1, "hi", asdf]
}
