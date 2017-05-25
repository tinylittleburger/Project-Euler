package utilities

import (
	"math/big"
	"strconv"
)

func IsPandigital(n int) bool {
	m := make([]bool, 10)
	l := 0

	for n > 0 {
		digit := n % 10
		n /= 10
		if m[digit] == true || digit == 0 {
			return false
		}

		m[digit] = true
		l++
	}

	for i := 1; i <= l; i++ {
		if m[i] == false {
			return false
		}
	}

	return true
}

func IsPalindrome(n int) bool {
	s := strconv.Itoa(n)

	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func IsPalindromeBig(n big.Int) bool {
	s := n.String()

	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
