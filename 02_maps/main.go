package main

import "fmt"

func main() {
	bar := map[string]float64{
		"amazon": 1700.0,
		"apple":  2000.0,
		"google": 1500.0,
	}

	// get a value
	fmt.Println(bar["amazon"])
	fmt.Println(bar["car"])

	value, ok := bar["car"]

	if !ok {
		fmt.Println("car is not present")
	} else {
		fmt.Println(value)
	}

	bar["car"] = 1400.0
	fmt.Println(bar)

	delete(bar, "car")
	fmt.Println(bar)

	for key := range bar {
		fmt.Println(key)
	}

	for key, value := range bar {
		fmt.Printf("%s - %.2f\n", key, value)
	}
}
