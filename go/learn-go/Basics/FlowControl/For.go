package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// The init and post statements are optional
	sumO := 1
	for sumO < 1000 {
		sumO += sumO
	}
	fmt.Println(sumO)

	// For is Go's while
	while := 1
	for while < 1000 {
		while += while
	}
	fmt.Println(while)

	// // loop forever
	// for {
	// 	fmt.Println("loop forever")
	// }
}
