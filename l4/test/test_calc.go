package test

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func matchString(a, b string) (bool, error) {
	return a == b, nil
}

type calcSign string

const (
	signAdd      calcSign = "+"
	signSubtract calcSign = "-"
	signMultiply calcSign = "*"
	signDivide   calcSign = "/"
)

const precision = 1e-10

type Operation interface {
	Values() (float64, float64)
	Solution(float64)
}

type operation struct {
	a      float64
	b      float64
	result float64
}

func (o operation) Values() (float64, float64) {
	return o.a, o.b
}

func (o operation) Solution(result float64) {
	o.result = result
}

func Calc(calc func(add, sub, mul, div <-chan Operation, ready chan<- bool)) {
	testCalc := func(t *testing.T) {
		tests := []struct {
			a    float64
			b    float64
			sign calcSign
			want float64
		}{
			{
				sign: signAdd,
				a:    10,
				b:    15,
				want: 25,
			},
			{
				sign: signSubtract,
				a:    3,
				b:    4,
				want: -1,
			},
			{
				sign: signMultiply,
				a:    6,
				b:    7,
				want: 42,
			},
			{
				sign: signDivide,
				a:    10,
				b:    3,
				want: 1.33333333333,
			},
		}
		add := make(chan Operation, 10)
		sub := make(chan Operation, 10)
		mul := make(chan Operation, 10)
		div := make(chan Operation, 10)
		ready := make(chan bool)
		go calc(add, sub, mul, div, ready)
		for _, tt := range tests {
			name := fmt.Sprintf("Checking calc for %s, %f, %f", tt.sign, tt.a, tt.b)
			t.Run(name, func(t *testing.T) {
				op := operation{
					a:      tt.a,
					b:      tt.b,
					result: tt.want,
				}
				switch tt.sign {
				case signAdd:
					add <- op
				case signSubtract:
					sub <- op
				case signMultiply:
					mul <- op
				case signDivide:
					div <- op
				}
				select {
				case <-ready:
					if math.Abs(op.result-tt.want) > precision {
						t.Errorf("Expected: %v, got: %v", tt.want, op.result)
					}
				case <-time.After(time.Second):
					t.Errorf("Waiting for result too long")
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
