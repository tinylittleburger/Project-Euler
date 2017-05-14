package main

import (
	"fmt"
)

func main() {
	size := 20
	mSize := size + 1
	m := make([][]int64, mSize)

	for i := 0; i < mSize; i++ {
		m[i] = make([]int64, mSize)
		m[i][0] = 1
		m[0][i] = 1
	}

	for i := 1; i < mSize; i++ {
		for j := 1; j < i; j++ {
			m[i][j] = m[i-1][j] + m[i][j-1]
			m[j][i] = m[i][j]
		}

		m[i][i] = m[i-1][i] + m[i][i-1]
	}

	fmt.Println(m[size][size])
}
