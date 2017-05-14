package main

import "comi/Project-Euler/utilities"
import "fmt"

func dividers(n int, sieve *utilities.Sieve) map[int]int {
	divs := make(map[int]int)

	i := 2
	for n > 1 {
		if sieve.IsPrime(i) && n%i == 0 {
			divs[i]++
			n /= i
		} else {
			i++
		}
	}

	return divs
}

func initialize() (map[int]int, map[int]int, map[int]int) {
	return make(map[int]int), make(map[int]int), make(map[int]int)
}

func main() {
	sieve := utilities.NewSieve()
	sieve.IsPrime(1000000)

	divs1, divs2, divs3 := initialize()

	for i := 2; ; i++ {
		var divs4 map[int]int

		if !sieve.IsPrime(i) {
			divs4 = dividers(i, &sieve)

			if len(divs1) == 4 && len(divs2) == 4 && len(divs3) == 4 && len(divs4) == 4 {
				fmt.Println(i - 3)
				return
			}
		} else {
			divs1, divs2, divs3 = initialize()
			divs4 = make(map[int]int)
		}

		divs1 = divs2
		divs2 = divs3
		divs3 = divs4
	}
}
