package main

import (
	"fmt"
)

func lastTenDigits(n, pow int) int {
	a := make([]int, 10)

	a[len(a)-1] = 1

	for i := 0; i < pow; i++ {
		multiply(a, n)
	}

	res := 0
	for _, v := range a {
		res = res*10 + v
	}

	return res
}

func multiply(a []int, n int) {
	carry := 0

	for i := len(a) - 1; i >= 0; i-- {
		temp := (a[i]*n + carry) / 10
		a[i] = (a[i]*n + carry) % 10
		carry = temp
	}
}

func main() {
	sum := int64(0)

	for i := 1; i <= 1000; i++ {
		sum += int64(lastTenDigits(i, i))
	}

	fmt.Println(sum % 10000000000)
}
