package main

import "fmt"

func split(sum int) (x, y int, w string) {
	x = sum * 4 / 9
	y = sum - x
	w = "123gogo"
	return
}

func main() {
	fmt.Println(split(17))
}
