package calc_test

import (
	"testing"

	"github.com/vegarsti/calc"
)

func TestExpressionValue(t *testing.T) {
	tcs := []struct {
		desc          string
		expression    calc.Expression
		expectedValue int
	}{
		{
			desc:          "1",
			expression:    calc.NewAtom(1),
			expectedValue: 1,
		},
		{
			desc:          "-1",
			expression:    calc.NewAtom(-1),
			expectedValue: -1,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.expression.Value()
			if got != tc.expectedValue {
				t.Errorf("expected %d, but got %d", tc.expectedValue, got)
			}
		})
	}
}
