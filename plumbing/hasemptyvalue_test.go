package plumbing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/internal"
)

func Test_HasEmptyValue(t *testing.T) {
	type testCase struct {
		Name          string
		Node          *goyaml.Node
		Expected      bool
		ErrorContains string
	}
	for _, tc := range []testCase{
		testCase{
			Name:     "TestCase1",
			Node:     internal.TestParseYAML(t, `"str"`),
			Expected: false,
		},
		testCase{
			Name:     "TestCase2",
			Node:     internal.TestParseYAML(t, `""`),
			Expected: true,
		},
		testCase{
			Name:     "TestCase3",
			Node:     internal.TestParseYAML(t, `false`),
			Expected: true,
		},
		testCase{
			Name:     "TestCase4",
			Node:     internal.TestParseYAML(t, `true`),
			Expected: false,
		},
		testCase{
			Name: "TestCase5",
			Node: &goyaml.Node{
				Kind:  goyaml.ScalarNode,
				Tag:   "!!bool",
				Value: "bla",
			},
			ErrorContains: "invalid node",
		},
		testCase{
			Name:     "TestCase6",
			Node:     internal.TestParseYAML(t, `null`),
			Expected: true,
		},
		testCase{
			Name:     "TestCase7",
			Node:     internal.TestParseYAML(t, `1`),
			Expected: false,
		},
		testCase{
			Name:     "TestCase8",
			Node:     internal.TestParseYAML(t, `0`),
			Expected: true,
		},
		testCase{
			Name: "TestCase9",
			Node: &goyaml.Node{
				Kind:  goyaml.ScalarNode,
				Tag:   "!!int",
				Value: "00",
			},
			ErrorContains: "invalid node",
		},
		testCase{
			Name: "TestCase10",
			Node: &goyaml.Node{
				Alias: &goyaml.Node{},
				Kind:  goyaml.AliasNode,
			},
			Expected: false,
		},
		testCase{
			Name:     "TestCase11",
			Node:     internal.TestParseYAML(t, `[]`),
			Expected: true,
		},
		testCase{
			Name:     "TestCase12",
			Node:     internal.TestParseYAML(t, `[1]`),
			Expected: false,
		},
		testCase{
			Name:     "TestCase13",
			Node:     internal.TestParseYAML(t, `{}`),
			Expected: true,
		},
		testCase{
			Name:     "TestCase14",
			Node:     internal.TestParseYAML(t, `{1:1}`),
			Expected: false,
		},
		testCase{
			Name:     "TestCase15",
			Node:     internal.TestParseYAML(t, `1.0`),
			Expected: false,
		},
		testCase{
			Name:     "TestCase16",
			Node:     internal.TestParseYAML(t, `0.0`),
			Expected: true,
		},
		testCase{
			Name:     "TestCase17",
			Node:     internal.TestParseYAML(t, `.nan`),
			Expected: false,
		},
		testCase{
			Name: "TestCase19",
			Node: &goyaml.Node{
				Kind: goyaml.ScalarNode,
				Tag:  "!!garbage",
			},
			ErrorContains: "invalid node",
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actual, err := HasEmptyValue(tc.Node)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tc.Expected, actual)
			}
		})
	}
}
