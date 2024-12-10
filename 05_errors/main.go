package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("sqrt of negative number: %v", n)
	}
	return math.Sqrt(n), nil

}

func main() {
	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Sqrt of 2.0 is: %f\n", s1)
	}

	s2, err := sqrt(-2.0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Sqrt of 2.0 is: %f\n", s2)
	}
}
