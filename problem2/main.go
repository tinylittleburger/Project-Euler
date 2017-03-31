package main

import (
	"fmt"
)

func main() {
	var sum int64
	limit := 4000000
	fib1 := 1
	fib2 := 2

	for fib2 < limit {
		if fib2%2 == 0 {
			sum += int64(fib2)
		}

		temp := fib1 + fib2
		fib1 = fib2
		fib2 = temp
	}

	fmt.Println(sum)
}
