package test

import (
	"fmt"
	"math"
	"testing"
)

// CalcOperation is a type that define allowed operations for calculator
type CalcOperation string

// Predefined operations
const (
	OperationAdd      = "+"
	OperationSubtract = "-"
	OperationMultiply = "*"
	OperationDivide   = "/"
)

const precision = 1e-10

// Calc is a function to test a function that returns operation function for operation
func Calc(calc func(operation CalcOperation) func(first, second float64) float64) {
	testCalc := func(t *testing.T) {
		tests := []struct {
			operation CalcOperation
			first     float64
			second    float64
			want      float64
		}{
			{
				operation: OperationAdd,
				first:     10,
				second:    15,
				want:      25,
			},
			{
				operation: OperationSubtract,
				first:     3,
				second:    4,
				want:      -1,
			},
			{
				operation: OperationMultiply,
				first:     6,
				second:    7,
				want:      42,
			},
			{
				operation: OperationDivide,
				first:     10,
				second:    3,
				want:      3.33333333333,
			},
		}
		for _, tt := range tests {
			name := fmt.Sprintf("Checking calc for %s, %f, %f", tt.operation, tt.first, tt.second)
			t.Run(name, func(t *testing.T) {
				t.Log()
				operation := calc(tt.operation)
				got := operation(tt.first, tt.second)
				if math.Abs(tt.want-got) > precision {
					t.Errorf("Expected: %v, got: %v", tt.want, got)
				}
			})
		}
	}
	testSuite := []testing.InternalTest{
		{
			Name: "TestCalc",
			F:    testCalc,
		},
	}

	testing.Main(matchString, testSuite, nil, nil)
}
