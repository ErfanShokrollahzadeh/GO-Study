package main

import "fmt"

func main() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("foo: %v type of %T\n", loons, loons)

	//lenght
	fmt.Println(len(loons))
	// index
	fmt.Println(loons[1])
	// slice
	fmt.Println(loons[1:])

	// for loop
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	// for value range
	for index := range loons {
		fmt.Println(index)
	}

	// for double value range
	for index, value := range loons {
		fmt.Printf("%s - %d\n", value, index)
	}
}
