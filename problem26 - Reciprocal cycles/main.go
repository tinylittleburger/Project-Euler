package main

import (
	"fmt"
)

func divCycle(divisor int) int {
	m := make(map[int]int)
	dividend := 10

	for index := 0; ; index++ {
		if dividend == 0 {
			return 0
		}

		if dividend < divisor {
			dividend *= 10
		} else if lastindex, ok := m[dividend]; !ok {
			m[dividend] = index
			dividend = (dividend % divisor) * 10
		} else {
			return index - lastindex
		}
	}
}

func main() {
	maxCycle := 0
	n := 0

	for i := 2; i < 1000; i++ {
		c := divCycle(i)
		if c > maxCycle {
			maxCycle = c
			n = i
		}
	}

	fmt.Println(n)
}
