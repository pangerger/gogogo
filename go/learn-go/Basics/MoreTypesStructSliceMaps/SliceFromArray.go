package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// note: changing the elements of a slie modifies the corresponding elements of its underlying array.
	// other slices that share the same underlying array will see the changes
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
