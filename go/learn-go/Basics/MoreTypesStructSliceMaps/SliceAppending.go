package main

import "fmt"

// append may create new array if the original array is too small
// And append() will impact to the original array
func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	tt := [2]int{1, 2}
	fmt.Printf("original array: len=%d cap=%d %v\n", len(tt), cap(tt), tt)

	ttS := tt[0:1]
	printSlice(ttS)

	ttS = append(ttS, 3)

	printSlice(ttS)

	fmt.Printf("original array: len=%d cap=%d %v\n", len(tt), cap(tt), tt)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
