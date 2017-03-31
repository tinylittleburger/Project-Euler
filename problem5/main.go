package main

import (
	"fmt"
)

func main() {
	limit := 20
	number := 1

	factors := make([]int, limit)

	for i := 0; i < limit-1; i++ {
		factors[i] = i + 2
	}

	i := 0

	for i < limit-1 {
		factor := factors[i]

		if factor != 0 {
			number *= factor
			j := i + 1

			for j < limit-1 {
				current := factors[j]
				if current%factor == 0 {
					if current == factor {
						factors[j] = 0
					} else {
						factors[j] = current / factor
					}
				}
				j++
			}
		}
		i++
	}

	fmt.Println(number)
}
