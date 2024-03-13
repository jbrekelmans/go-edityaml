package plumbing

import (
	"fmt"
	"strings"

	goyaml "gopkg.in/yaml.v3"
)

// nodeKind is a helper type for converting kinds to a string (for debugging / logging / error messages).
type nodeKind goyaml.Kind

func (n nodeKind) String() string {
	kind := goyaml.Kind(n)
	var sb strings.Builder
	append := func(s string) {
		if sb.Len() > 0 {
			sb.WriteString("|")
		}
		sb.WriteString(s)
	}
	if (kind & goyaml.DocumentNode) != 0 {
		kind &= ^goyaml.DocumentNode
		append("document")
	}
	if (kind & goyaml.SequenceNode) != 0 {
		kind &= ^goyaml.SequenceNode
		append("sequence")
	}
	if (kind & goyaml.MappingNode) != 0 {
		kind &= ^goyaml.MappingNode
		append("mapping")
	}
	if (kind & goyaml.ScalarNode) != 0 {
		kind &= ^goyaml.ScalarNode
		append("scalar")
	}
	if (kind & goyaml.AliasNode) != 0 {
		kind &= ^goyaml.AliasNode
		append("alias")
	}
	if kind != 0 {
		return fmt.Sprintf("0x%X", kind)
	} else if sb.Len() == 0 {
		return "0"
	}
	return sb.String()
}
