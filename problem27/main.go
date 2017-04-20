package main

import (
	"fmt"
)

func sieve(a []bool) {
	for i := 2; i < len(a); i++ {
		if !a[i] {
			for j := 2; i*j < len(a); j++ {
				a[i*j] = true
			}
		}
	}
}

func quadratic(n int, a int, b int) int {
	return n*n + a*n + b
}

func countPrimes(a int, b int, notprime []bool) int {
	count := 0

	for i := 0; ; i++ {
		q := quadratic(i, a, b)

		if q >= len(notprime) {
			break
		}

		if q > 0 && !notprime[q] {
			count++
		}
	}

	return count
}

func main() {
	notprime := make([]bool, 10000)
	sieve(notprime)
	limit := 1000
	count := 0
	var a, b int

	for i := -limit + 1; i < limit; i++ {
		for j := -limit; j <= limit; j++ {
			c := countPrimes(i, j, notprime)

			if c > count {
				a = i
				b = j
				count = c
			}
		}
	}

	fmt.Printf("a:%d b:%d count:%d", a, b, count)

}
