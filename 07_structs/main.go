package main

import "fmt"

type Trade struct {
	Symbol string
	Valume int
	Price  float64
	buy    bool
}

func main() {
	t1 := Trade{"apple", 10, 99.98, true}
	fmt.Println(t1)
	fmt.Printf("%+v\n", t1)
	fmt.Println(t1.Symbol)
	fmt.Println(t1.Valume)
}
