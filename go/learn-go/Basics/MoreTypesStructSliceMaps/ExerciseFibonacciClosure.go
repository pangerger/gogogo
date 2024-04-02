package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var i, j int = -1, 0

	return func() int {

		result := i + j
		if result < 0 {
			result = 0
			i = j
			j = 1
		}

		if j == 0 {
			i = 0
			j = 1
		} else {
			i = j
			j = result
		}

		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
