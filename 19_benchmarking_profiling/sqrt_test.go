package main

import (
	"fmt"
	"testing"
)

func almostEqual(a, b float64) bool {
	return Abs(a - b) < 0.001
}

func TestAbs(t *testing.T) {
	val , err := Sqrt(2)

	if err != nil {
		t.Fatalf("Error in Sqrt(2)")
	}

	if !almostEqual(val, 1.414) {
		t.Fatalf("Expected approximately 1.414, but got %v", val)
	}
} 


type testcase struct {
	value float64
	expected float64
}

func TestMany(t *testing.T){
	tasecases := []testcase{
		{0,0},
		{2.0, 1.414},
		{9.0, 3.0},
	}

	for _, tc := range tasecases {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T){
			out, err := Sqrt(tc.value)
			if err != nil {
				t.Fatalf("Error in Sqrt(%f)", tc.value)
			}

			if !almostEqual(out, tc.expected) {
				t.Fatalf("Expected approximately %f, but got %v", tc.expected, out)
			}
			
		})
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Sqrt(float64(i))
		if err != nil {
			b.Fatalf("Error in Sqrt(%f)", float64(i))
		}
	}
}