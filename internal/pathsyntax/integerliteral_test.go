package pathsyntax

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jbrekelmans/go-edityaml/internal"
)

func Test_parseIntegerLiteral(t *testing.T) {
	type testCase struct {
		Name             string
		Text             string
		ErrorContains    string
		ExpectedTextRest string
		ExpectedValue    *big.Int
	}
	for _, tc := range []testCase{
		testCase{
			Name:          "TestCase1",
			Text:          `00`,
			ErrorContains: `leading zeros`,
		},
		testCase{
			Name:             "TestCase2",
			Text:             `0`,
			ExpectedTextRest: "",
			ExpectedValue:    internal.TestParseBigInt(t, "0"),
		},
		testCase{
			Name:             "TestCase3",
			Text:             `0a`,
			ExpectedTextRest: "a",
			ExpectedValue:    internal.TestParseBigInt(t, "0"),
		},
		testCase{
			Name:          "TestCase4",
			Text:          `-`,
			ErrorContains: `unexpected character "-"`,
		},
		testCase{
			Name:          "TestCase5",
			Text:          `-a`,
			ErrorContains: `unexpected character "-"`,
		},
		testCase{
			Name:          "TestCase6",
			Text:          `-0`,
			ErrorContains: `leading zeros or signed zero`,
		},
		testCase{
			Name:             "TestCase7",
			Text:             `-1`,
			ExpectedTextRest: "",
			ExpectedValue:    internal.TestParseBigInt(t, "-1"),
		},
		testCase{
			Name:             "TestCase8",
			Text:             `-1a`,
			ExpectedTextRest: "a",
			ExpectedValue:    internal.TestParseBigInt(t, "-1"),
		},
		testCase{
			Name:             "TestCase9",
			Text:             `2147483647b`,
			ExpectedTextRest: "b",
			ExpectedValue:    internal.TestParseBigInt(t, "2147483647"),
		},
		testCase{
			Name:             "TestCase10",
			Text:             `-2147483648d`,
			ExpectedTextRest: "d",
			ExpectedValue:    internal.TestParseBigInt(t, "-2147483648"),
		},
		testCase{
			Name:          "TestCase11",
			Text:          `1234`,
			ExpectedValue: internal.TestParseBigInt(t, "1234"),
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actualValue, actualTextRest, err := parseIntegerLiteral(tc.Text)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tc.ExpectedTextRest, actualTextRest)
				assert.Equal(t, tc.ExpectedValue, actualValue)
			}
		})
	}
}
