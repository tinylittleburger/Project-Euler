package main

import (
	"fmt"
	"math"
)

var powers []int

func digitPowers(n int) []int {
	a := make([]int, 10)

	for i := range a {
		a[i] = int(math.Pow(float64(i), float64(n)))
	}

	return a
}

func powDigitSum(n int) int {
	sum := 0

	for n > 0 {
		digit := n % 10
		n = n / 10
		sum += powers[digit]
	}

	return sum
}

func main() {
	powers = digitPowers(5)
	sum := 0

	for i := 10; i < 1000000; i++ {
		if i == powDigitSum(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
