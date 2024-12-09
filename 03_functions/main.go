package main

import "fmt"

func add(a int, b int) int {
	return a + b

}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	val := add(1, 2)
	println(val)

	div, mod := divide(7, 2)
	fmt.Printf("div=%d mod=%d\n", div, mod)
}
