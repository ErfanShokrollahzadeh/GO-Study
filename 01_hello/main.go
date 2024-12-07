package main

import "fmt"

func main() {
	book := "The Colour of Magic"

	fmt.Println(book)
	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	fmt.Println(book[4:11])
	fmt.Println(book[4:])
	fmt.Println(book[:4])

	fmt.Println("hi " + book)

	// multiline string
	poem := `
The road goes ever on
Down from the door where it began
Now far ahead the road has gone
...
`
	fmt.Println(poem)
}
