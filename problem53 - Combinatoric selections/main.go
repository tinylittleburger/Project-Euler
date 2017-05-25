package main

import (
	"fmt"
)

func main() {
	size := 100
	limit := 1000000
	count := 0
	m := make([][]int, size+1)

	for i := 0; i < len(m); i++ {
		m[i] = make([]int, size+1)
		m[i][0] = 1

		if i > 0 {
			for j := 1; j < len(m[i]); j++ {
				m[i][j] = m[i-1][j-1] + m[i-1][j]

				if m[i][j] > limit {
					m[i][j] = limit
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
