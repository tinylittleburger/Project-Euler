package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	return n == reverseInt(n)
}

func isPalindromeAlt(n int) bool {
	s := strconv.Itoa(n)

	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func reverseInt(number int) int {
	reversed := 0

	for number > 0 {
		digit := number % 10
		number /= 10
		reversed = reversed*10 + digit
	}

	return reversed
}

func main() {
	max := 1

	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			number := i * j
			if isPalindrome(number) && number > max {
				max = number
			}
		}
	}

	fmt.Printf("%d", max)
}
