package main

import "comi/Project-Euler/utilities"
import "fmt"

func primeFactors(n int, sieve *utilities.Sieve) map[int]int {
	factors := make(map[int]int)
	i := 2

	for n > 1 {
		if sieve.IsPrime(i) && n%i == 0 {
			factors[i]++
			n /= i
		} else {
			i++
		}
	}

	return factors
}

func main() {
	sieve := utilities.NewSieve()
	sieve.IsPrime(1000000)

	n := 4
	factors := make([]map[int]int, n-1, n)

	for i := 2; ; i++ {
		if !sieve.IsPrime(i) {
			factors = append(factors, primeFactors(i, &sieve))
			found := true

			for _, v := range factors {
				if len(v) != n {
					found = false
					break
				}
			}

			if found {
				fmt.Println(i - n + 1)
				return
			}
		} else {
			factors = make([]map[int]int, n-1, n)
		}

		factors = factors[1:n:n]
	}
}
