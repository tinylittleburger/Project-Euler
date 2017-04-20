package main

import (
	"fmt"
	"math"
)

func main() {
	primes := 0
	var n float64 = 2
	limit := 10001
	var wantedPrime int

	for {
		prime := true

		for i := 2; i <= int(math.Sqrt(n)); i++ {
			if int(n)%i == 0 {
				prime = false
				break
			}
		}

		if prime {
			primes++
			if primes == limit {
				wantedPrime = int(n)
				break
			}
		}
		n++
	}

	fmt.Println(wantedPrime)
}
