package pathsyntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseStringLiteral(t *testing.T) {
	type testCase struct {
		Name             string
		Text             string
		ErrorContains    string
		ExpectedTextRest string
		ExpectedValue    string
	}
	for _, tc := range []testCase{
		testCase{
			Name:          "TestCase1",
			Text:          `"`,
			ErrorContains: `unterminated string literal`,
		},
		testCase{
			Name:          "TestCase2",
			Text:          `"\`,
			ErrorContains: `unterminated string literal`,
		},
		testCase{
			Name:             "TestCase3",
			Text:             `""`,
			ExpectedTextRest: "",
			ExpectedValue:    "",
		},
		testCase{
			Name:             "TestCase4",
			Text:             `"\\".remainder`,
			ExpectedTextRest: ".remainder",
			ExpectedValue:    "\\",
		},
		testCase{
			Name:             "TestCase5",
			Text:             "\"\x80\\\\\".remainder",
			ExpectedTextRest: ".remainder",
			ExpectedValue:    "\uFFFD\\",
		},
		testCase{
			Name:             "TestCase6",
			Text:             "\"\x80\".remainder",
			ExpectedTextRest: ".remainder",
			ExpectedValue:    "\uFFFD",
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actualValue, actualTextRest, err := parseStringLiteral(tc.Text)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tc.ExpectedTextRest, actualTextRest)
				assert.Equal(t, tc.ExpectedValue, actualValue)
			}
		})
	}
}
