package main

import (
	"fmt"
	"math"
)

func penthagon(n int) int {
	return n * (3*n - 1) / 2
}

func main() {
	m := make(map[int]int)
	minD := math.MaxInt32

	for i := 1; i < 10000; i++ {
		n := penthagon(i)
		m[n]++
	}

	for a := range m {
		for b := range m {
			if a != b {
				if _, ok1 := m[a+b]; ok1 {
					d := int(math.Abs(float64(a - b)))
					if _, ok2 := m[d]; ok2 && d < minD {
						minD = d
					}
				}
			}
		}
	}

	fmt.Println(minD)

}
