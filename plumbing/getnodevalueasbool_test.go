package plumbing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/internal"
)

func Test_getNodeValueAsBool(t *testing.T) {
	type testCase struct {
		Name          string
		Node          *goyaml.Node
		Expected      bool
		ErrorContains string
	}
	for _, tc := range []testCase{
		testCase{
			Name:     "TestCase1",
			Node:     internal.TestParseYAML(t, `true`),
			Expected: true,
		},
		testCase{
			Name:     "TestCase2",
			Node:     internal.TestParseYAML(t, `false`),
			Expected: false,
		},
		testCase{
			Name: "TestCase3",
			Node: &goyaml.Node{
				Kind:  goyaml.ScalarNode,
				Tag:   "!!bool",
				Value: "badValue",
			},
			ErrorContains: `invalid node`,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actual, err := getNodeValueAsBool(tc.Node)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tc.Expected, actual)
			}
		})
	}
}
