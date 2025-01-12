package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol can't be empty")
	}
	if volume <= 0 {
		return nil, fmt.Errorf("volume must be greater than zero")
	}
	if price <= 0.0 {
		return nil, fmt.Errorf("price must be greater than zero")
	}
	return &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}, nil
}

func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {
	t, err := NewTrade("MSFT", 10, 99.98, true)
	if err != nil {
		fmt.Printf("error: can't create trade - %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", t.Value())
}
