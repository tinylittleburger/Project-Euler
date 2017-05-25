package main

import (
	"comi/Project-Euler/utilities"
	"fmt"
	"math"
	"math/big"
)

func isLychrel(n int) bool {
	return isLychrelInternal(int(n), 50)
}

func isLychrelInternal(n int, count int) bool {
	if count == 0 {
		return true
	}

	if n > math.MaxInt32/10 {
		return isLychrelInternalBig(big.NewInt(int64(n)), count)
	}

	nRev := utilities.ReverseInt(n)
	newN := nRev + n

	if utilities.IsPalindrome(newN) {
		return false
	}
	return isLychrelInternal(newN, count-1)
}

func isLychrelInternalBig(n *big.Int, count int) bool {
	if count == 0 {
		return true
	}

	nRev := utilities.ReverseString(n.String())
	nRevBig, _ := big.NewInt(0).SetString(nRev, 10)
	newN := n.Add(n, nRevBig)

	if utilities.IsPalindromeBig(*newN) {
		return false
	}

	return isLychrelInternalBig(newN, count-1)
}

func main() {
	lychrel := 0

	for i := 10; i < 10000; i++ {
		if isLychrel(i) {
			lychrel++
		}
	}

	fmt.Println(lychrel)
}
