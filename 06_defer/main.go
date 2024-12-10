package main

import "fmt"

func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {
	defer cleanup("worker")
	fmt.Println("Worker started")
}

func main() {
	worker()
}
