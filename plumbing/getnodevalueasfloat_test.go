package plumbing

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/internal"
)

func Test_getNodeValueAsFloat(t *testing.T) {
	type testCase struct {
		Name          string
		Node          *goyaml.Node
		Expected      float64
		ErrorContains string
	}
	for _, tc := range []testCase{
		testCase{
			Name:     "TestCase1",
			Node:     internal.TestParseYAML(t, `.NAN`),
			Expected: math.NaN(),
		},
		testCase{
			Name:     "TestCase2",
			Node:     internal.TestParseYAML(t, `.NaN`),
			Expected: math.NaN(),
		},
		testCase{
			Name:     "TestCase3",
			Node:     internal.TestParseYAML(t, `.nan`),
			Expected: math.NaN(),
		},
		testCase{
			Name:     "TestCase4",
			Node:     internal.TestParseYAML(t, `+.inf`),
			Expected: math.Inf(1),
		},
		testCase{
			Name:     "TestCase5",
			Node:     internal.TestParseYAML(t, `.inf`),
			Expected: math.Inf(1),
		},
		testCase{
			Name:     "TestCase6",
			Node:     internal.TestParseYAML(t, `.Inf`),
			Expected: math.Inf(1),
		},
		testCase{
			Name:     "TestCase7",
			Node:     internal.TestParseYAML(t, `0.0`),
			Expected: 0.0,
		},
		testCase{
			Name:     "TestCase8",
			Node:     internal.TestParseYAML(t, `-.inf`),
			Expected: math.Inf(-1),
		},
		testCase{
			Name:     "TestCase9",
			Node:     internal.TestParseYAML(t, `1.75e2`),
			Expected: 175.0,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actual, err := getNodeValueAsFloat(tc.Node)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				if math.IsNaN(tc.Expected) {
					assert.Truef(t, math.IsNaN(actual), "expected NaN")
				} else {
					assert.Equal(t, tc.Expected, actual)
				}
			}
		})
	}
}
