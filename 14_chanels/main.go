package main

// we have 2 kinds of channels: unbuffered and buffered

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// <- ch // this will block the program
	// fmt.Println("this will not be printed")

	go func() {
		// send number of the chanel
		ch <- 350
	}()

	// receive number from the chanel
	val := <-ch
	fmt.Printf("got %d\n", val)
	fmt.Println("===============")

	// buffered channel (multiple values)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(1*time.Second)
		}
		// close the channel
	} ()

	// receive number from the chanel
	for i := 0; i < 5; i++ {
		val := <-ch
		fmt.Printf("receved %d\n", val)
	}

}
