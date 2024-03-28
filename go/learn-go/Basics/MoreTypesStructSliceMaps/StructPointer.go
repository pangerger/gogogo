package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// (*p).X can simple write as p.X if p is a pointer to Sturct
func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
