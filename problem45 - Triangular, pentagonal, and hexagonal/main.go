package main

import (
	"fmt"
)

func triangle(n int) int {
	return n * (n + 1) / 2
}

func pentagonal(n int) int {
	return n * (3*n - 1) / 2
}

func hexagonal(n int) int {
	return n * (2*n - 1)
}

func main() {
	tris := make(map[int]int)
	pents := make(map[int]int)
	hexs := make(map[int]int)

	for i := 1; i <= 100000; i++ {
		tris[triangle(i)]++
		pents[pentagonal(i)]++
		hexs[hexagonal(i)]++
	}

	for tri := range tris {
		if _, ok1 := pents[tri]; ok1 {
			if _, ok2 := hexs[tri]; ok2 {
				fmt.Println(tri)
			}
		}
	}
}
