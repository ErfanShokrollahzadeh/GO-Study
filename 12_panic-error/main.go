package main

import "fmt"


func main(){
	vals := []int{1,2,3}
	// panic: runtime error: index out of range [3] with length 3
	v := vals[10]
	fmt.Println(v)
}