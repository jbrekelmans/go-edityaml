package pathsyntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isDigit(t *testing.T) {
	type testCase struct {
		Name     string
		Input    byte
		Expected bool
	}
	for _, tc := range []testCase{
		testCase{
			Name:     "TestCase1",
			Input:    'a',
			Expected: false,
		},
		testCase{
			Name:     "TestCase2",
			Input:    '\x00',
			Expected: false,
		},
		testCase{
			Name:     "TestCase3",
			Input:    '0',
			Expected: true,
		},
		testCase{
			Name:     "TestCase4",
			Input:    '9',
			Expected: true,
		},
		testCase{
			Name:     "TestCase5",
			Input:    '~',
			Expected: false,
		},
		testCase{
			Name:     "TestCase6",
			Input:    'A',
			Expected: false,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actual := isDigit(tc.Input)
			assert.Equal(t, tc.Expected, actual)
		})
	}
}
