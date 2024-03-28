package main

import "fmt"

// Deferred function calls are pushed onto a stack
// When a function returns, its deferred calls are executed in last-in-first-out order
// reade this blog for more information: https://go.dev/blog/defer-panic-and-recover
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	fmt.Println("ready for return")
}
