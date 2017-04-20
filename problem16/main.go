package main

import (
	"fmt"
)

func main() {
	pow := 1000
	a := []int{2}

	for i := 2; i <= pow; i++ {
		carry := 0

		for j := 0; j < len(a); j++ {
			temp := a[j]*2 + carry
			carry = temp / 10
			a[j] = temp % 10
		}

		if carry != 0 {
			a = append(a, carry)
		}
	}

	sum := 0
	for j := 0; j < len(a); j++ {
		sum += a[j]
	}

	fmt.Println(sum)
}
