package main

import (
	"comi/Project-Euler/utilities"
	"fmt"
)

func rotations(a []int) [][]int {
	res := [][]int{}

	for i := 1; i < len(a); i++ {
		res = append(res, rotate(a, i))
	}

	return res
}

func rotate(a []int, index int) []int {
	return append(a[index:], a[:index]...)
}

func toArray(i int) []int {
	var a []int
	for i > 0 {
		a = append(a, i%10)
		i /= 10
	}

	var res []int
	for i := len(a) - 1; i >= 0; i-- {
		res = append(res, a[i])
	}

	return res
}

func toNumber(a []int) int {
	n := 0

	for i := 0; i < len(a); i++ {
		n *= 10
		n += a[i]
	}

	return n
}

func main() {
	sieve := utilities.NewSieve()
	count := 0

	for i := 2; i < 1000000; i++ {
		if sieve.IsPrime(i) {
			rots := rotations(toArray(i))
			prime := true

			for j := range rots {
				if !sieve.IsPrime(toNumber(rots[j])) {
					prime = false
					break
				}
			}

			if prime {
				count++
			}
		}
	}

	fmt.Println(count)
}
