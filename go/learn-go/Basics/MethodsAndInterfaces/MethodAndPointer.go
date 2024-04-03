package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// value for type
func (v Vertex) Abs() float64 {
	fmt.Println("x: %d, y: %d", v.X, v.Y)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer for type
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) ScaleValueCopy(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	v.ScaleValueCopy(10)
	v.Abs()
}
