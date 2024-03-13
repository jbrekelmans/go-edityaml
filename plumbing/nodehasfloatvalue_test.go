package plumbing

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/internal"
)

func Test_nodeHasFloatValue(t *testing.T) {
	type testCase struct {
		Name          string
		Node          *goyaml.Node
		Value         any
		Expected      bool
		ErrorContains string
	}
	for _, tc := range []testCase{
		testCase{
			Name: "TestCase1",
			Node: &goyaml.Node{
				Kind:  goyaml.ScalarNode,
				Tag:   "!!float",
				Value: "badValue",
			},
			ErrorContains: "invalid node",
		},
		testCase{
			Name:     "TestCase2",
			Node:     internal.TestParseYAML(t, `1.0`),
			Value:    "not-a-float",
			Expected: false,
		},
		testCase{
			Name:     "TestCase3",
			Node:     internal.TestParseYAML(t, `1.0`),
			Value:    1.0,
			Expected: true,
		},
		testCase{
			Name:     "TestCase4",
			Node:     internal.TestParseYAML(t, `.nan`),
			Value:    math.NaN(),
			Expected: true,
		},
		testCase{
			Name:     "TestCase5",
			Node:     internal.TestParseYAML(t, `.nan`),
			Value:    1.0,
			Expected: false,
		},
		testCase{
			Name:     "TestCase6",
			Node:     internal.TestParseYAML(t, `1.0`),
			Value:    "not-a-float",
			Expected: false,
		},
		testCase{
			Name:     "TestCase7",
			Node:     internal.TestParseYAML(t, `1.0`),
			Value:    float32(1.0),
			Expected: true,
		},
		testCase{
			Name:     "TestCase8",
			Node:     internal.TestParseYAML(t, `.nan`),
			Value:    float32(math.NaN()),
			Expected: true,
		},
		testCase{
			Name:     "TestCase9",
			Node:     internal.TestParseYAML(t, `.nan`),
			Value:    float32(1.0),
			Expected: false,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actual, err := nodeHasFloatValue(tc.Node, tc.Value)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tc.Expected, actual)
			}
		})
	}
}
