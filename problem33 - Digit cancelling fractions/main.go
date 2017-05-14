package main

import (
	"fmt"
)

func isCurious(a, b int) bool {
	if a == b {
		return false
	}

	digA1 := float64(a % 10)
	digA2 := float64(a / 10)
	digB1 := float64(b % 10)
	digB2 := float64(b / 10)

	var div2 float64
	switch {
	case digA1 == digB1 && digA1 != 0:
		div2 = digA2 / digB2
	case digA1 == digB2 && digA1 != 0:
		div2 = digA2 / digB1
	case digA2 == digB1 && digA2 != 0:
		div2 = digA1 / digB2
	case digA2 == digB2 && digA2 != 0:
		div2 = digA1 / digB1
	default:
		return false
	}

	div1 := float64(a) / float64(b)

	return div1 == div2
}

func canonizeFraction(a, b int) (int, int) {
	limit := a

	for i := 2; i*i <= limit; i++ {
		for a%i == 0 && b%i == 0 {
			a /= i
			b /= i
		}
	}

	return a, b
}

func main() {
	a := 1
	b := 1

	for i := 10; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			if isCurious(i, j) {
				a *= i
				b *= j
			}
		}
	}

	a1, b1 := canonizeFraction(a, b)
	fmt.Printf("%d %d\n", a1, b1)
}
