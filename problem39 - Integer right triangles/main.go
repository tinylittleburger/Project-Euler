package main

import (
	"fmt"
)

func isRATriangle(a, b, c int) bool {
	return c*c == a*a+b*b
}

func main() {
	m := make(map[int]int)
	limit := 1000

	for i := 1; i <= limit-2; i++ {
		for j := 1; i+j <= limit-1; j++ {
			for k := 1; k <= 1000-i-j; k++ {
				if isRATriangle(i, j, k) {
					m[i+j+k]++
				}
			}
		}
	}

	maxValue := 0
	maxKey := 0

	for key, value := range m {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}
	fmt.Println(maxKey)
}
