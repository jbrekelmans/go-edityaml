package plumbing

import (
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jbrekelmans/go-edityaml/internal"
)

func Test_convertBigIntToInt(t *testing.T) {
	type testCase struct {
		Name          string
		Int           *big.Int
		Expected      int
		ErrorContains string
	}
	for _, tc := range []testCase{
		testCase{
			Name:          "TestCase1",
			Int:           internal.TestParseBigInt(t, "9223372036854775808"), // int64 max value + 1
			ErrorContains: "out of range of int",
		},
		testCase{
			Name:          "TestCase2",
			Int:           internal.TestParseBigInt(t, "-9223372036854775809"), // int64 min value - 1
			ErrorContains: "out of range of int",
		},
		testCase{
			Name:     "TestCase3",
			Int:      internal.TestParseBigInt(t, "1234"),
			Expected: 1234,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actual, err := convertBigIntToInt(tc.Int)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tc.Expected, actual)
			}
		})
	}
}

func Test_convertInt64ToInt(t *testing.T) {
	// TODO test on architectures where int and int64 are not the same size.
	actual, ok := convertInt64ToInt(1234)
	if assert.True(t, ok) {
		assert.Equal(t, 1234, actual)
	}
}

func Test_convertUint32ToInt(t *testing.T) {
	// TODO test on architectures where int maxValue fits in uint32.
	actual, err := convertUint32ToInt(1234)
	if assert.NoError(t, err) {
		assert.Equal(t, 1234, actual)
	}
}

func Test_convertToInt(t *testing.T) {
	type testCase struct {
		Name          string
		Value         any
		Expected      int
		ErrorContains string
	}
	for _, tc := range []testCase{
		// TODO test on architectures with different uint/int sizes
		testCase{
			Name:     "TestCase1",
			Value:    int8(1),
			Expected: 1,
		},
		testCase{
			Name:     "TestCase2",
			Value:    int16(1),
			Expected: 1,
		},
		testCase{
			Name:     "TestCase3",
			Value:    int32(1),
			Expected: 1,
		},
		testCase{
			Name:     "TestCase4",
			Value:    int64(1),
			Expected: 1,
		},
		testCase{
			Name:     "TestCase5",
			Value:    1,
			Expected: 1,
		},
		testCase{
			Name:     "TestCase6",
			Value:    uint8(1),
			Expected: 1,
		},
		testCase{
			Name:     "TestCase7",
			Value:    uint16(1),
			Expected: 1,
		},
		testCase{
			Name:     "TestCase8",
			Value:    uint32(1),
			Expected: 1,
		},
		testCase{
			Name:          "TestCase9",
			Value:         uint64(0xFFFF_FFFF_FFFF_FFFF),
			ErrorContains: "uint64 value out of range of int",
		},
		testCase{
			Name:     "TestCase10",
			Value:    uint64(1),
			Expected: 1,
		},
		testCase{
			Name:          "TestCase11",
			Value:         uint(math.MaxUint),
			ErrorContains: "uint value out of range of int",
		},
		testCase{
			Name:     "TestCase12",
			Value:    uint(1),
			Expected: 1,
		},
		testCase{
			Name:     "TestCase13",
			Value:    *internal.TestParseBigInt(t, "1"),
			Expected: 1,
		},
		testCase{
			Name:     "TestCase14",
			Value:    internal.TestParseBigInt(t, "1"),
			Expected: 1,
		},
		testCase{
			Name:          "TestCase15",
			Value:         (*big.Int)(nil),
			ErrorContains: "cannot convert nil",
		},
		testCase{
			Name:          "TestCase16",
			Value:         "str",
			ErrorContains: "value is not an integer type",
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			actual, err := convertToInt(tc.Value)
			if tc.ErrorContains != "" {
				assert.ErrorContains(t, err, tc.ErrorContains)
			} else if assert.NoError(t, err) {
				assert.Equal(t, tc.Expected, actual)
			}
		})
	}
}
