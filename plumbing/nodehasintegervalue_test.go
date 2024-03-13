package plumbing

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	goyaml "gopkg.in/yaml.v3"

	"github.com/jbrekelmans/go-edityaml/internal"
)

func Test_nodeHasIntegerValue(t *testing.T) {
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
			Node:     internal.TestParseYAML(t, `1`),
			Value:    int8(1),
			Expected: true,
		},
		testCase{
			Name:     "TestCase2",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    int8(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase3",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    int16(1),
			Expected: true,
		},
		testCase{
			Name:     "TestCase4",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    int16(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase5",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    int32(1),
			Expected: true,
		},
		testCase{
			Name:     "TestCase6",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    int32(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase7",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    int64(1),
			Expected: true,
		},
		testCase{
			Name:     "TestCase8",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    int64(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase9",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    int(1),
			Expected: true,
		},
		testCase{
			Name:     "TestCase10",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    int(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase11",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    uint8(1),
			Expected: true,
		},
		testCase{
			Name:     "TestCase12",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    uint8(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase13",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    uint16(1),
			Expected: true,
		},
		testCase{
			Name:     "TestCase14",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    uint16(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase15",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    uint32(1),
			Expected: true,
		},
		testCase{
			Name:     "TestCase16",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    uint32(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase17",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    uint64(1),
			Expected: true,
		},
		testCase{
			Name:     "TestCase18",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    uint64(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase19",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    uint(1),
			Expected: true,
		},
		testCase{
			Name:     "TestCase20",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    uint(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase21",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    *big.NewInt(1),
			Expected: true,
		},
		testCase{
			Name:     "TestCase22",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    *big.NewInt(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase23",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    big.NewInt(1),
			Expected: true,
		},
		testCase{
			Name:     "TestCase24",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    big.NewInt(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase25",
			Node:     internal.TestParseYAML(t, `1`),
			Value:    nil,
			Expected: false,
		},
		testCase{
			Name: "TestCase26",
			Node: &goyaml.Node{
				Kind:  goyaml.ScalarNode,
				Tag:   "!!int",
				Value: "00",
			},
			ErrorContains: "invalid node",
		},
		testCase{
			Name: "TestCase27",
			Node: &goyaml.Node{
				Kind:  goyaml.ScalarNode,
				Tag:   "!!int",
				Value: "9223372036854775807000", // int64 maxValue * 10
			},
			Value:    1,
			Expected: false,
		},
		testCase{
			Name: "TestCase28",
			Node: &goyaml.Node{
				Kind:  goyaml.ScalarNode,
				Tag:   "!!int",
				Value: "9223372036854775807000", // int64 maxValue * 10
			},
			Value:    uint(1),
			Expected: false,
		},
		testCase{
			Name:     "TestCase29",
			Node:     internal.TestParseYAML(t, `2`),
			Value:    (*big.Int)(nil),
			Expected: false,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actual, err := nodeHasIntegerValue(tc.Node, tc.Value)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tc.Expected, actual)
			}
		})
	}
}
