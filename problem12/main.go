package main

import "fmt"

func factors(number int64) int {
	factors := 0
	i := int64(1)

	for ; i*i < number; i++ {
		if number%i == 0 {
			factors = factors + 2
		}
	}

	if i*i == number {
		factors++
	}

	return factors
}

func main() {
	triangle := int64(0)
	for i := 1; ; i++ {
		triangle += int64(i)

		factors := factors(triangle)
		if factors > 500 {
			fmt.Println(triangle)
			return
		}
	}
}
