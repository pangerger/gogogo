package main

import (
	"fmt"
	"math"
)

// z -= (z*z - x) / (2*z)
// (Note: If you are interested in the details of the algorithm, the z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. This general approach is called Newton's method. It works well for many functions but especially well for square root.)
func Sqrt(x float64) float64 {
	diff := 0.001

	z := 1.0

	var cnt int = 10
	for i := 0; i < cnt; i++ {
		temp := z*z - x
		if temp < 0 {
			temp = -temp
		}
		if temp < diff {
			fmt.Println("caculation times is", i)
			return z
		}
		z -= (z*z - x) / (2 * z)
	}

	fmt.Println("caculation times is", cnt)
	return z
}

func main() {
	var n float64 = 2.0
	fmt.Println(Sqrt(n))
	fmt.Println(math.Sqrt(n))
}
