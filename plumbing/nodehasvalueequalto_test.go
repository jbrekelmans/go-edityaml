package plumbing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/internal"
)

func Test_NodeHasValueEqualTo(t *testing.T) {
	type testCase struct {
		Name          string
		Node          *goyaml.Node
		Value         any
		Expected      bool
		ErrorContains string
	}
	for _, tc := range []testCase{
		testCase{
			Name:     "TestCase1",
			Node:     internal.TestParseYAML(t, `value`),
			Value:    "value",
			Expected: true,
		},
		testCase{
			Name:     "TestCase2",
			Node:     internal.TestParseYAML(t, `false`),
			Value:    false,
			Expected: true,
		},
		testCase{
			Name:     "TestCase3",
			Node:     internal.TestParseYAML(t, `true`),
			Value:    false,
			Expected: false,
		},
		testCase{
			Name: "TestCase4",
			Node: &goyaml.Node{
				Kind:  goyaml.ScalarNode,
				Tag:   "!!bool",
				Value: "bla",
			},
			Value:         false,
			ErrorContains: "invalid node",
		},
		testCase{
			Name:     "TestCase5",
			Node:     internal.TestParseYAML(t, `null`),
			Value:    nil,
			Expected: true,
		},
		testCase{
			Name:     "TestCase6",
			Node:     internal.TestParseYAML(t, `null`),
			Value:    &struct{}{},
			Expected: false,
		},
		testCase{
			Name:     "TestCase7",
			Node:     internal.TestParseYAML(t, `null`),
			Value:    (*struct{})(nil),
			Expected: true,
		},
		testCase{
			Name:     "TestCase8",
			Node:     internal.TestParseYAML(t, `null`),
			Value:    struct{}{},
			Expected: false,
		},
		testCase{
			Name:     "TestCase9",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    1,
			Expected: true,
		},
		testCase{
			Name:     "TestCase10",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    -1,
			Expected: false,
		},
		testCase{
			Name: "TestCase11",
			Node: &goyaml.Node{
				Kind:  goyaml.ScalarNode,
				Tag:   "!!int",
				Value: "00",
			},
			ErrorContains: "invalid node",
		},
		testCase{
			Name: "TestCase12",
			Node: func() *goyaml.Node {
				node := internal.TestParseYAML(t, `1`)
				return &goyaml.Node{
					Alias: node,
					Kind:  goyaml.AliasNode,
				}
			}(),
			Value:    1,
			Expected: true,
		},
		testCase{
			Name: "TestCase13",
			Node: func() *goyaml.Node {
				node := &goyaml.Node{
					Kind: goyaml.AliasNode,
				}
				node.Alias = node
				return node
			}(),
			ErrorContains: "aborting after resolving",
		},
		testCase{
			Name: "TestCase14",
			Node: &goyaml.Node{
				Kind: goyaml.MappingNode,
				Tag:  "!!map",
			},
			ErrorContains: "NodeHasValueEqualTo is not implemented",
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actual, err := NodeHasValueEqualTo(tc.Node, tc.Value)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tc.Expected, actual)
			}
		})
	}
}
