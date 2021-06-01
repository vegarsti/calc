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
		{
			desc: "1+1",
			expression: calc.NewBinaryExpression(
				calc.NewAtom(1),
				calc.NewAtom(1),
				calc.Plus,
			),
			expectedValue: 2,
		},
		{
			desc: "1-1",
			expression: calc.NewBinaryExpression(
				calc.NewAtom(1),
				calc.NewAtom(1),
				calc.Minus,
			),
			expectedValue: 0,
		},
		{
			desc: "2*2",
			expression: calc.NewBinaryExpression(
				calc.NewAtom(2),
				calc.NewAtom(2),
				calc.Multiply,
			),
			expectedValue: 4,
		},
		{
			desc: "1/2",
			expression: calc.NewBinaryExpression(
				calc.NewAtom(1),
				calc.NewAtom(2),
				calc.Divide,
			),
			expectedValue: 0,
		},
		{
			desc: "2/1",
			expression: calc.NewBinaryExpression(
				calc.NewAtom(2),
				calc.NewAtom(1),
				calc.Divide,
			),
			expectedValue: 2,
		},
		{
			desc: "2 ^ 2",
			expression: calc.NewBinaryExpression(
				calc.NewAtom(2),
				calc.NewAtom(2),
				calc.Exponentiate,
			),
			expectedValue: 4,
		},
		{
			desc: "2 ^ 3",
			expression: calc.NewBinaryExpression(
				calc.NewAtom(2),
				calc.NewAtom(3),
				calc.Exponentiate,
			),
			expectedValue: 8,
		},
		{
			desc: "(1+1)+(1+1)",
			expression: calc.NewBinaryExpression(
				calc.NewBinaryExpression(
					calc.NewAtom(1),
					calc.NewAtom(1),
					calc.Plus,
				),
				calc.NewBinaryExpression(
					calc.NewAtom(1),
					calc.NewAtom(1),
					calc.Plus,
				),
				calc.Plus,
			),
			expectedValue: 4,
		},
		{
			desc: "(1+1)+(1+1)",
			expression: calc.NewBinaryExpression(
				calc.NewBinaryExpression(
					calc.NewAtom(1),
					calc.NewAtom(1),
					calc.Plus,
				),
				calc.NewBinaryExpression(
					calc.NewAtom(1),
					calc.NewAtom(1),
					calc.Plus,
				),
				calc.Minus,
			),
			expectedValue: 0,
		},
		{
			desc: "(1+1)/(1+1)",
			expression: calc.NewBinaryExpression(
				calc.NewBinaryExpression(
					calc.NewAtom(1),
					calc.NewAtom(1),
					calc.Plus,
				),
				calc.NewBinaryExpression(
					calc.NewAtom(1),
					calc.NewAtom(1),
					calc.Plus,
				),
				calc.Divide,
			),
			expectedValue: 1,
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
