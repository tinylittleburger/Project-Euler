package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(s string) bool {
	for i := 0; i*2 < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	sum := 0

	for i := 1; i < 1000000; i++ {
		a := strconv.Itoa(i)
		b := strconv.FormatInt(int64(i), 2)

		if isPalindrome(a) && isPalindrome(b) {
			sum += i
		}
	}

	fmt.Println(sum)
}
