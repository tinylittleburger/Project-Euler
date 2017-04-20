package main

import (
	"fmt"
)

func chainLength(number int, terms []int, limit int) int {
	if number == 1 {
		terms[1] = 1
		return 1
	}

	if number < limit && terms[number] != 0 {
		return terms[number]
	}

	length := 0
	if number%2 == 0 {
		length = chainLength(number/2, terms, limit) + 1
	} else {
		length = chainLength(number*3+1, terms, limit) + 1
	}

	if number < limit {
		terms[number] = length
	}
	return length
}

func main() {
	limit := 1000000
	terms := make([]int, limit)
	max := 1
	maxPosition := 1

	for i := 1; i < limit; i++ {
		length := chainLength(i, terms, limit)

		if length > max {
			maxPosition = i
			max = length
		}
	}

	fmt.Printf("Chain length: %d, for number: %d", maxPosition, max)
}
