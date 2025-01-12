package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

type Circle struct {
	Radius float64
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type shape interface {
	Area() float64
}

func sumAreas(shapes []shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

func main() {
	s := Square{10}
	c := Circle{5}

	totalArea := sumAreas([]shape{&s, &c})
	fmt.Println(totalArea)
}
