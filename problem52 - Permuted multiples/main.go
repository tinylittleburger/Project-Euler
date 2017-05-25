package main

import (
	"fmt"
)

func isPermutablyCurious(n int, a []int) bool {
	m := extractDigits(n)

	for _, v := range a {
		if !areMapsEqual(m, extractDigits(n*v)) {
			return false
		}
	}

	return true
}

func areMapsEqual(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func extractDigits(n int) map[int]int {
	m := make(map[int]int)

	for n > 0 {
		m[n%10]++
		n /= 10
	}

	return m
}

func main() {
	a := []int{2, 3, 4, 5, 6}

	for i := 1; ; i++ {
		if isPermutablyCurious(i, a) {
			fmt.Println(i)
			return
		}
	}
}
