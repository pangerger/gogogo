package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		result[i] = make([]uint8, dx, dx)
		for j := 0; j < dx; j++ {
			result[i][j] = uint8(i * j)
		}
	}
	return result
}

func main() {
	result := Pic(8, 10)
	fmt.Printf("len=%d, cap=%d, content=%v", len(result), cap(result), result)
}
