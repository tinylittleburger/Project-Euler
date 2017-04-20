package main

import (
	"fmt"
)

func main() {
	fib1 := []int{1}
	fib2 := []int{1}
	index := 2
	digits := 1

	for {
		carry := 0

		for j := 0; j < len(fib2); j++ {
			temp := fib2[j] + fib1[j] + carry
			carry = temp / 10
			temp2 := fib2[j]
			fib2[j] = temp % 10
			fib1[j] = temp2
		}
		index++

		if carry != 0 {
			fib1 = append(fib1, 0)
			fib2 = append(fib2, carry)
			digits++

			if digits >= 1000 {
				break
			}
		}
	}

	fmt.Println(index)
}
