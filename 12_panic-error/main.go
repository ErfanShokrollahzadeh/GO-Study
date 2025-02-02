package main

import (
	"fmt"
	// "os"
)


func safevalue(vals []int, index int) int {
	defer func () {
		if err:= recover(); err != nil {
			fmt.Printf("panic occured: %v\n", err)
		}
	}()

	return vals[index] 
		
}



func main(){
	// vals := []int{1,2,3}
	// // panic: runtime error: index out of range [3] with length 3
	// v := vals[10]
	// fmt.Println(v)

	// file, err := os.Open("no-such-file")

	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()
	// fmt.Println("file opened")


	v := safevalue([]int{1,2,3}, 10)
	fmt.Println(v)
}