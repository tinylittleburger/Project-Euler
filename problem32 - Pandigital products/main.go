package main

import (
	"fmt"
)

func isPan(a, b, c int) bool {
	m := make([]bool, 10)

	for a > 0 {
		digit := a % 10
		a /= 10
		if m[digit] == true || digit == 0 {
			return false
		}

		m[digit] = true
	}

	for b > 0 {
		digit := b % 10
		b /= 10
		if m[digit] == true || digit == 0 {
			return false
		}

		m[digit] = true
	}

	for c > 0 {
		digit := c % 10
		c /= 10
		if m[digit] == true || digit == 0 {
			return false
		}

		m[digit] = true
	}

	for i, value := range m {
		if i != 0 && value == false {
			return false
		}
	}

	return true
}

func getLimits(n int) (int, int) {
	if n < 10 {
		return 1000, 99999
	}

	return 100, 9999
}

func main() {
	pans := make(map[int]int)

	for i := 1; i <= 99; i++ {
		from, to := getLimits(i)

		for j := from; j <= to; j++ {
			product := i * j

			if isPan(i, j, product) {
				pans[product]++
			}
		}
	}

	sum := 0
	for value := range pans {
		sum += value
	}

	fmt.Println(sum)
}
