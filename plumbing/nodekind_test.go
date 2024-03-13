package plumbing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	goyaml "gopkg.in/yaml.v3"
)

func Test_nodeKindToString(t *testing.T) {
	type testCase struct {
		Kind     goyaml.Kind
		Expected string
		Name     string
	}
	for _, tc := range []testCase{
		testCase{
			Kind:     goyaml.ScalarNode,
			Expected: "scalar",
			Name:     "TestCase1",
		},
		testCase{
			Kind:     0,
			Expected: "0",
			Name:     "TestCase2",
		},
		testCase{
			Kind:     goyaml.MappingNode,
			Expected: "mapping",
			Name:     "TestCase3",
		},
		testCase{
			Kind:     goyaml.SequenceNode,
			Expected: "sequence",
			Name:     "TestCase4",
		},
		testCase{
			Kind:     goyaml.DocumentNode,
			Expected: "document",
			Name:     "TestCase5",
		},
		testCase{
			Kind:     goyaml.AliasNode,
			Expected: "alias",
			Name:     "TestCase6",
		},
		testCase{
			Kind:     goyaml.Kind(uint32(3) << 30),
			Expected: "0xC0000000",
			Name:     "TestCase7",
		},
		testCase{
			Kind:     goyaml.MappingNode | goyaml.SequenceNode,
			Expected: "sequence|mapping",
			Name:     "TestCase8",
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actual := nodeKind(tc.Kind).String()
			assert.Equal(t, tc.Expected, actual)
		})
	}
}
