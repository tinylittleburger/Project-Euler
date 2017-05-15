package main

import "comi/Project-Euler/utilities"
import "fmt"

func main() {
	sieve := utilities.NewSieve()

	sieve.IsPrime(10000)

	for i := 3; ; i += 2 {
		if !sieve.IsPrime(i) {
			goldbach := false

			for j := 1; j <= i; j++ {
				k := i - 2*j*j

				if sieve.IsPrime(k) {
					goldbach = true
					break
				}
			}

			if !goldbach {
				fmt.Println(i)
				break
			}
		}
	}
}
