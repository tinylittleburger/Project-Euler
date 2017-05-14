package main

import "comi/Project-Euler/utilities"
import "fmt"

func main() {
	sieve := utilities.NewSieve()

	for i := 987654321; i > 0; i-- {
		if sieve.IsPrime(i) && utilities.IsPandigital(i) {
			fmt.Println(i)
			break
		}
	}
}
