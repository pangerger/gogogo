package main

import "fmt"

// length: The length of a slice is the number of elemnets it contains.
// capacity: The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[0:0]
	printSlice(s)

	// Extend its length.
	s = s[3:4]
	printSlice(s)

	// Drop its first two values.
	s = s[0:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
