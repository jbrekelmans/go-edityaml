package internal

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsNilSafe(t *testing.T) {
	type testCase struct {
		Input         reflect.Value
		ExpectedIsNil bool
		ExpectedOK    bool
		Name          string
	}
	for _, tc := range []testCase{
		testCase{
			Input:         reflect.ValueOf((*struct{})(nil)),
			ExpectedIsNil: true,
			ExpectedOK:    true,
			Name:          "TestCase1",
		},
		testCase{
			Input:         reflect.ValueOf(&struct{}{}),
			ExpectedIsNil: false,
			ExpectedOK:    true,
			Name:          "TestCase2",
		},
		testCase{
			Input:         reflect.ValueOf(struct{}{}),
			ExpectedIsNil: false,
			ExpectedOK:    false,
			Name:          "TestCase3",
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			isNil, ok := IsNilSafe(tc.Input)
			assert.Equalf(t, tc.ExpectedIsNil, isNil, "isNil")
			assert.Equalf(t, tc.ExpectedOK, ok, "ok")
		})
	}
}
