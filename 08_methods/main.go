package main

import "fmt"

type Trade struct {
	Symbol string
	Valume int
	Price  float64
	buy    bool
}

// Add a method to the Trade struct that returns the value of the trade.

func (t *Trade) Value() float64 {
	value := float64(t.Valume) * t.Price
	if t.buy {
		value = -value
	}

	return value
}

func main() {
	t := Trade{"apple", 10, 99.98, true}
	fmt.Println(t.Value())
}
