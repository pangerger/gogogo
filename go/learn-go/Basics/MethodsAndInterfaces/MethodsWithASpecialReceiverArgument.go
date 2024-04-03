package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// function with struct
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	w := Vertex{3, 4}
	fmt.Println(w.Abs())
}
