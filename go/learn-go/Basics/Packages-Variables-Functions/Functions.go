package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addOmitTypeButTheLast(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(addOmitTypeButTheLast(12, 21))
}
