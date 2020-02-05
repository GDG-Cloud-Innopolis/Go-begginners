package test

import (
	"flag"
	"math"
	"testing"
)

func matchString(a, b string) (bool, error) {
	return a == b, nil
}

const precision = 1e-10

// TestEquation is a function to test a square equation solver
func TestEquation(EquationRoots func(a, b, c float64) (float64, float64)) {
	testEquation := func(t *testing.T) {
		flag.Set("test.v", "")
		type args struct {
			a float64
			b float64
			c float64
		}
		tests := []struct {
			name  string
			args  args
			want  float64
			want1 float64
		}{
			{
				name: "6x² + 11x - 35 = 0",
				args: args{
					6,
					11,
					-35,
				},
				want:  -3.5,
				want1: 5. / 3.,
			},
			{
				name: "2x² - 4x - 2 = 0",
				args: args{
					2,
					-4,
					-2,
				},
				want:  1 - math.Sqrt(2),
				want1: 1 + math.Sqrt(2),
			},
			{
				name: "2x² - 64 = 0",
				args: args{
					2,
					0,
					-64,
				},
				want:  -4 * math.Sqrt(2),
				want1: 4 * math.Sqrt(2),
			},
			{
				name: "x² - 16 = 0",
				args: args{
					1,
					0,
					-16,
				},
				want:  -4,
				want1: 4,
			},
			{
				name: "x² - 7x = 0",
				args: args{
					1,
					-7,
					0,
				},
				want:  0,
				want1: 7,
			},
			{
				name: "2x² + 8x = 0",
				args: args{
					2,
					8,
					0,
				},
				want:  -4,
				want1: 0,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, got1 := EquationRoots(tt.args.a, tt.args.b, tt.args.c)
				t.Logf("Equation: %s\nRoots: %f, %f\nYour answer: %f, %f", tt.name, tt.want, tt.want1, got, got1)
				if !(math.Abs(got-tt.want) < precision) {
					t.Errorf("Incorrect x1")
				}
				if !(math.Abs(got1-tt.want1) < precision) {
					t.Errorf("Incorrect x2")
				}
			})
		}
	}
	testSuite := []testing.InternalTest{
		{
			Name: "TestEquation",
			F:    testEquation,
		},
	}

	flag.Set("test.v", "=true")
	testing.Init()
	testing.Main(matchString, testSuite, nil, nil)
}
