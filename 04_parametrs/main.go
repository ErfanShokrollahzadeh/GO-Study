package main

import "fmt"

func doubleAT(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

func main() {
	values := []int{1, 2, 3, 4}
	doubleAT(values, 2)
	fmt.Println(values)
	// Output: [1 2 6 4]

	val := 10
	double(val)
	fmt.Println(val)
}
