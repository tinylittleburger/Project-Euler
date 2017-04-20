package main

import "math/big"
import "fmt"

func main() {
	limit := 100
	fact := big.NewInt(1)

	for i := 1; i <= limit; i++ {
		fact.Mul(fact, big.NewInt(int64(i)))
	}

	zero := big.NewInt(0)
	ten := big.NewInt(10)
	sum := big.NewInt(0)

	for fact.Cmp(zero) > 0 {
		sum.Add(big.NewInt(0).Mod(fact, ten), sum)
		fact.Div(fact, ten)
	}

	fmt.Println(sum)
}
