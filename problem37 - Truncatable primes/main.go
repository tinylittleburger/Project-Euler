package main

import (
	"comi/Project-Euler/utilities"
	"fmt"
)

func truncates(n int) []int {
	a := []int{}

	temp1 := n
	temp2 := n
	pow := 1

	for temp1 > 10 {
		temp1 /= 10
		a = append(a, temp1)
		pow *= 10
		a = append(a, temp2%pow)
	}

	return a
}

func isTruncPrime(sieve *utilities.Sieve, n int) bool {
	if !sieve.IsPrime(n) {
		return false
	}

	a := truncates(n)
	for _, value := range a {
		if !sieve.IsPrime(value) {
			return false
		}
	}

	return true
}

func main() {
	sieve := utilities.NewSieve()
	sum := 0
	count := 0

	for i := 1000000; i >= 11; i-- {
		if isTruncPrime(&sieve, i) {
			fmt.Println(i)
			sum += i
			count++
		}
	}

	fmt.Println("----------------")
	fmt.Println(count)
	fmt.Println(sum)
}
