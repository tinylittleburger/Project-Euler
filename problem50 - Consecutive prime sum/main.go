package main

import "comi/Project-Euler/utilities"
import "fmt"

func main() {
	var primes []int
	sieve := utilities.NewSieve()
	limit := 1000000
	sieve.Init(10000000)

	for i := 2; i < limit*2; i++ {
		if sieve.IsPrime(i) {
			primes = append(primes, i)
		}
	}

	longest := 0
	num := 0

	for i := range primes {
		length := 0
		current := 0

		for j := i; current < limit; j++ {
			length++
			current += primes[j]

			if sieve.IsPrime(current) && length > longest {
				longest = length
				num = current
			}
		}
	}

	fmt.Println(num)
	fmt.Println(longest)
}
