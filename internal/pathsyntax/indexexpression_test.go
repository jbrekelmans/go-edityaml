package pathsyntax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jbrekelmans/go-edityaml/internal"
)

func Test_ParseIndexExpression(t *testing.T) {
	type testCase struct {
		Name             string
		Arg              string
		ErrorContains    string
		ExpectedPathRest string
		ExpectedPathItem any
	}
	for _, tc := range []testCase{
		testCase{
			Name:          "TestCase1",
			Arg:           `[`,
			ErrorContains: `unterminated index expression`,
		},
		testCase{
			Name:          "TestCase2",
			Arg:           `["`,
			ErrorContains: `unterminated string literal`,
		},
		testCase{
			Name:          "TestCase3",
			Arg:           `[""`,
			ErrorContains: `unterminated index expression`,
		},
		testCase{
			Name:          "TestCase4",
			Arg:           `[""a`,
			ErrorContains: `string/integer literal in an index expression must be immediately followed by "]" character`,
		},
		testCase{
			Name:             "TestCase5",
			Arg:              `[""]a`,
			ExpectedPathRest: "a",
			ExpectedPathItem: "",
		},
		testCase{
			Name:             "TestCase6",
			Arg:              `["a"]b`,
			ExpectedPathRest: "b",
			ExpectedPathItem: "a",
		},
		testCase{
			Name:          "TestCase7",
			Arg:           `[00]`,
			ErrorContains: `invalid integer literal`,
		},
		testCase{
			Name:             "TestCase8",
			Arg:              `[0]`,
			ExpectedPathItem: internal.TestParseBigInt(t, "0"),
		},
		testCase{
			Name:             "TestCase9",
			Arg:              `[2]c`,
			ExpectedPathItem: internal.TestParseBigInt(t, "2"),
			ExpectedPathRest: "c",
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actualPathItem, actualPathRest, err := ParseIndexExpression(tc.Arg)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tc.ExpectedPathItem, actualPathItem)
				assert.Equal(t, tc.ExpectedPathRest, actualPathRest)
			}
		})
	}
}
