package main

import "comi/Project-Euler/utilities"
import "fmt"

func main() {
	sieve := utilities.NewSieve()

	sieve.IsPrime(10000)

	for i := 2; ; i++ {
		if i%2 != 0 && !(sieve.IsPrime(i)) {
			goldbach := false

			for j := 1; j <= i; j++ {
				k := i - 2*j*j

				if sieve.IsPrime(k) {
					goldbach = true
				}
			}

			if !goldbach {
				fmt.Println(i)
				break
			}
		}
	}
}
