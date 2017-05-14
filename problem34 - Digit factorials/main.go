package main

import (
	"fmt"
)

var digitFactorials = getDigitFactorials()

func getDigitFactorials() []int {
	a := make([]int, 10)
	a[0] = 1

	for i := 1; i < len(a); i++ {
		a[i] = i * a[i-1]
	}

	return a
}

func isCurious(n int) bool {
	sum := 0
	digits := n

	for digits > 0 {
		digit := digits % 10
		digits /= 10
		sum += digitFactorials[digit]
	}

	return n == sum
}

func main() {
	sum := 0

	for i := 3; i < 10000000; i++ {
		if isCurious(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
