package main

import (
	"fmt"
	"math/big"
)

func main() {
	size := 100
	a := make(map[string]int)
	for i := 2; i <= size; i++ {
		for j := 2; j <= size; j++ {
			powd := new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(int64(j)), nil).String()
			a[powd] = a[powd] + 1
		}
	}

	fmt.Println(len(a))
}
