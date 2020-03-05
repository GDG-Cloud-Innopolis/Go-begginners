//START ONE OMIT
package math

import (
	"testing"
)

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

// END ONE OMIT

// //START TWO OMIT

func TestAverage1(t *testing.T) {
	type args struct {
		xs []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Two numbers",
			args: args{
				xs: []float64{1, 2},
			},
			want: 1.5,
		},
		{
			name: "Equal numbers",
			args: args{
				xs: []float64{1, 1, 1, 1, 1, 1},
			},
			want: 1,
		},
	}
	// END TWO OMIT
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.xs); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

// END THREE OMIT

// START FOUR OMIT
func BenchmarkAverage10(b *testing.B) {
	// run the Average function b.N times
	for n := 0; n < b.N; n++ {
		Average([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

// END FOUR OMIT
