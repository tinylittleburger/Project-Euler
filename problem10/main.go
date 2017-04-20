package main

import (
	"fmt"
)

func main() {
	size := 2000000
	sieve := make([]bool, size)

	for i := 2; i < size; i++ {
		if !sieve[i] {
			for j := 2; i*j < size; j++ {
				sieve[i*j] = true
			}
		}
	}

	var sum int64

	for i := 2; i < size; i++ {
		if !sieve[i] {
			sum += int64(i)
		}
	}

	fmt.Println(sum)
}
