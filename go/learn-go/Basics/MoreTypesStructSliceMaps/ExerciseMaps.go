package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)

	var cArray []string = strings.Fields(s)

	for _, v := range cArray {
		m[v] = m[v] + 1
	}

	return m
}

func main() {
	result := WordCount("ok sir")
	fmt.Printf("content:", result)
}
