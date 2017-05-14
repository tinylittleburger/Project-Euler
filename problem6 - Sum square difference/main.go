package main

import (
	"fmt"
)

func main() {
	limit := 100

	sum := (limit + 1) * limit / 2
	sumSquared := sum * sum
	sumOfSquares := 0

	for i := 1; i <= limit; i++ {
		sumOfSquares += i * i
	}

	fmt.Println(sumSquared - sumOfSquares)
}
