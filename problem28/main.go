package main

import (
	"fmt"
)

func main() {
	size := 1001
	sum := 1
	current := 1
	step := 2

	for i := 1; i <= size/2; i++ {
		for j := 1; j <= 4; j++ {
			current += step
			sum += current
		}

		step += 2
	}

	fmt.Println(sum)
}
