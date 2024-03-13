package plumbing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	goyaml "gopkg.in/yaml.v3"
)

func Test_DoMapLookup(t *testing.T) {
	type testCase struct {
		Name          string
		Node          *goyaml.Node
		Key           any
		Expected      int
		ErrorContains string
	}
	for _, tc := range []testCase{
		testCase{
			Name:          "TestCase1",
			Node:          &goyaml.Node{},
			ErrorContains: "node is not a mapping",
		},
		testCase{
			Name: "TestCase2",
			Node: &goyaml.Node{
				Kind: goyaml.MappingNode,
				Content: []*goyaml.Node{
					&goyaml.Node{},
				},
			},
			ErrorContains: "odd number of children",
		},
		testCase{
			Name: "TestCase3",
			Node: &goyaml.Node{
				Kind: goyaml.MappingNode,
				Content: []*goyaml.Node{
					nil,
					&goyaml.Node{},
				},
			},
			ErrorContains: `Content[0] is nil`,
		},
		testCase{
			Name: "TestCase4",
			Node: &goyaml.Node{
				Kind: goyaml.MappingNode,
				Content: []*goyaml.Node{
					&goyaml.Node{
						Kind:  goyaml.ScalarNode,
						Tag:   "!!bool",
						Value: "huh?",
					},
					&goyaml.Node{},
				},
			},
			ErrorContains: `invalid node`,
		},
		testCase{
			Name: "TestCase5",
			Node: &goyaml.Node{
				Kind: goyaml.MappingNode,
				Content: []*goyaml.Node{
					&goyaml.Node{
						Kind:  goyaml.ScalarNode,
						Tag:   "!!bool",
						Value: "false",
					},
					nil,
				},
			},
			Key:           false,
			ErrorContains: `Content[1] is nil`,
		},
		testCase{
			Name: "TestCase6",
			Node: &goyaml.Node{
				Kind: goyaml.MappingNode,
				Content: []*goyaml.Node{
					&goyaml.Node{
						Kind:  goyaml.ScalarNode,
						Tag:   "!!bool",
						Value: "false",
					},
					&goyaml.Node{},
				},
			},
			Key:      false,
			Expected: 1,
		},
		testCase{
			Name: "TestCase7",
			Node: &goyaml.Node{
				Kind: goyaml.MappingNode,
				Content: []*goyaml.Node{
					&goyaml.Node{
						Kind:  goyaml.ScalarNode,
						Tag:   "!!bool",
						Value: "false",
					},
					&goyaml.Node{},
				},
			},
			Key:      true,
			Expected: -1,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actual, err := DoMapLookup(tc.Node, tc.Key)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tc.Expected, actual)
			}
		})
	}
}
