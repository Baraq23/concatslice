package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 6, 7, 8}
	fmt.Printf("Slice1: %v  Slice2: %v\n", slice1, slice2)
	fmt.Println("Concatenated: ", ConcatSlice(slice1, slice2))
}

func ConcatSlice(slice1, slice2 []int) []int {
	// The append function is used to concatenate slice2 to slice1
	// The ... operator is used to expand slice2 into a list of arguments for append
	s := append(slice1, slice2...)
	return s
}
