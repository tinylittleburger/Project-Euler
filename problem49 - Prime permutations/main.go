package main

import "comi/Project-Euler/utilities"
import "fmt"

func main() {
	sieve := utilities.NewSieve()
	found := 0

	for i := 1000; ; i++ {
		if sieve.IsPrime(i) {
			m := make(map[int]int)
			a := utilities.ToSlice(i)

			for ; a != nil; a = utilities.NextPermutation(a) {
				n := utilities.ToNumber(a)

				if i != n && sieve.IsPrime(n) {
					m[n-i]++
				}
			}

			for v := range m {
				if _, ok := m[v*2]; ok {
					found++

					if found == 2 {
						fmt.Println(i)
						fmt.Println(v)
						return
					}
				}
			}
		}
	}
}
