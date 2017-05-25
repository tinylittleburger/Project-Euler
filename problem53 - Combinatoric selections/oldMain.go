package main

import "math/big"
import "fmt"

var zero = big.NewInt(int64(0))
var decrease = big.NewInt(int64(-1))
var million = big.NewInt(int64(1000000))

func comb(n, r int) *big.Int {
	nFact := fact(n)
	rFact := fact(r)
	nrFact := fact(n - r)

	result := rFact.Mul(rFact, nrFact)
	result = result.Div(nFact, result)

	return result
}

func fact(n int) *big.Int {
	bigN := big.NewInt(int64(n))
	f := big.NewInt(int64(1))

	for r := bigN.Cmp(zero); r > 0; r = bigN.Cmp(zero) {
		f.Mul(f, bigN)
		bigN.Add(bigN, decrease)
	}

	return f
}

func main() {
	count := 0

	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			c := comb(n, r)

			if c.Cmp(million) >= 0 {
				fmt.Printf("%d_C_%d \n", n, r)
				count++
			}
		}
	}

	fmt.Println(count)
}
