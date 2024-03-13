package edityaml

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParsePath(t *testing.T) {
	type testCase struct {
		Name          string
		Arg           string
		ErrorContains string
		Expected      Path
	}
	for _, tc := range []testCase{
		testCase{
			Name:     "TestCase1",
			Arg:      ``,
			Expected: Path{},
		},
		testCase{
			Name:          "TestCase2",
			Arg:           `.0`,
			ErrorContains: `path is invalid`,
		},
		testCase{
			Name:     "TestCase3",
			Arg:      `.a.b`,
			Expected: Path{"a", "b"},
		},
		testCase{
			Name:     "TestCase4",
			Arg:      `.a["b"]`,
			Expected: Path{"a", "b"},
		},
		testCase{
			Name:          "TestCase5",
			Arg:           `["`,
			ErrorContains: `error parsing index expression`,
		},
		testCase{
			Name:     "TestCase6",
			Arg:      `["c"].d`,
			Expected: Path{"c", "d"},
		},
		testCase{
			Name:          "TestCase7",
			Arg:           `["a"]garbage`,
			ErrorContains: `path is invalid`,
		},
		testCase{
			Name:          "TestCase8",
			Arg:           `garbage`,
			ErrorContains: `path is invalid`,
		},
		testCase{
			Name:     "TestCase9",
			Arg:      `["a"]["\n"]`,
			Expected: Path{"a", "\n"},
		},
		testCase{
			Name:     "TestCase10",
			Arg:      `["a"][0]`,
			Expected: Path{"a", new(big.Int).SetInt64(0)},
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actual, err := ParsePath(tc.Arg)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tc.Expected, actual)
			}
		})
	}
}

func Test_Path(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		type testCase struct {
			Name     string
			Input    Path
			Expected string
		}
		for _, tc := range []testCase{
			testCase{
				Name:     "TestCase1",
				Input:    nil,
				Expected: "<empty path>",
			},
			testCase{
				Name:     "TestCase2",
				Input:    Path{},
				Expected: "<empty path>",
			},
			testCase{
				Name:     "TestCase3",
				Input:    Path{new(big.Int).SetInt64(1), "0", "abc"},
				Expected: `[1]["0"].abc`,
			},
			testCase{
				Name:     "TestCase4",
				Input:    Path{struct{}{}},
				Expected: ".<key of type struct {}>",
			},
			testCase{
				Name:     "TestCase5",
				Input:    Path{"a", new(big.Int).SetInt64(1), []any{"item"}},
				Expected: ".a[1].<key of type []interface {}>",
			},
			testCase{
				Name:     "TestCase6",
				Input:    Path{(*big.Int)(nil), nil, ([]string)(nil), struct{}{}},
				Expected: ".<null>.<null>.<null>.<key of type struct {}>",
			},
			testCase{
				Name:     "TestCase7",
				Input:    Path{int8(-1)},
				Expected: "[-1]",
			},
			testCase{
				Name:     "TestCase8",
				Input:    Path{int16(-1)},
				Expected: "[-1]",
			},
			testCase{
				Name:     "TestCase9",
				Input:    Path{int32(-1)},
				Expected: "[-1]",
			},
			testCase{
				Name:     "TestCase10",
				Input:    Path{-1},
				Expected: "[-1]",
			},
			testCase{
				Name:     "TestCase11",
				Input:    Path{uint8(2)},
				Expected: "[2]",
			},
			testCase{
				Name:     "TestCase12",
				Input:    Path{uint16(2)},
				Expected: "[2]",
			},
			testCase{
				Name:     "TestCase13",
				Input:    Path{uint32(2)},
				Expected: "[2]",
			},
			testCase{
				Name:     "TestCase14",
				Input:    Path{uint64(2)},
				Expected: "[2]",
			},
			testCase{
				Name:     "TestCase15",
				Input:    Path{uint(2)},
				Expected: "[2]",
			},
			testCase{
				Name:     "TestCase16",
				Input:    Path{*new(big.Int).SetInt64(10)},
				Expected: "[10]",
			},
		} {
			t.Run(tc.Name, func(t *testing.T) {
				actual := tc.Input.String()
				assert.Equal(t, tc.Expected, actual)
			})
		}
	})
}
