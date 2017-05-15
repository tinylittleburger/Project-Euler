package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := big.NewInt(0)
	m := big.NewInt(10000000000)

	for i := 1; i <= 1000; i++ {
		current := big.NewInt(int64(i))
		current.Exp(current, current, m)
		sum.Add(sum, current)
	}

	fmt.Println(sum.Mod(sum, m).String())
}
