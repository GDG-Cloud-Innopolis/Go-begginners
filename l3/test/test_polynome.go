package test

import (
	"math"
	"testing"
)

func matchString(a, b string) (bool, error) {
	return a == b, nil
}

// Polynomial is a function to test a polynomial function
func Polynomial(polynomial func([]float64) func(x float64) float64) {
	testPolynomial := func(t *testing.T) {
		tests := []struct {
			name    string
			polynom []float64
			x       float64
			want    float64
		}{
			{
				name:    "2x+1",
				polynom: []float64{2, 1},
				x:       10,
				want:    21,
			},
			{
				name:    "3x^2+2x+1",
				polynom: []float64{3, 2, 1},
				x:       15,
				want:    706,
			},
			{
				name:    "4x^3+3x^2+2x+1",
				polynom: []float64{4, 3, 2, 1},
				x:       20,
				want:    33241,
			},
			{
				name:    "5x^4+4x^3+3x^2+2x+1",
				polynom: []float64{5, 4, 3, 2, 1},
				x:       25,
				want:    2017551,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Log()
				solver := polynomial(tt.polynom)
				got := solver(tt.x)
				if math.Abs(tt.want-got) > precision {
					t.Errorf("Expected: %v, got: %v", tt.want, got)
				}
			})
		}
	}
	testSuite := []testing.InternalTest{
		{
			Name: "TestPolynomial",
			F:    testPolynomial,
		},
	}

	testing.Main(matchString, testSuite, nil, nil)
}
