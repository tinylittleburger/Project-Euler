package main

import (
	"fmt"
	"sort"
	"strconv"
)

func isPan(n int) bool {
	m := make([]bool, 10)

	for n > 0 {
		digit := n % 10
		n /= 10
		if m[digit] == true || digit == 0 {
			return false
		}

		m[digit] = true
	}

	for i, value := range m {
		if i != 0 && value == false {
			return false
		}
	}

	return true
}

func isPanMultiple(n int, p int) (bool, int) {
	s := ""
	for i := 1; i <= p; i++ {
		s += strconv.Itoa(n * i)
	}

	if len(s) != 9 {
		return false, 0
	}

	a, _ := strconv.Atoi(s)

	if isPan(a) {
		return true, a
	}

	return false, 0
}

func main() {
	a := []int{}

	for i := 1; i < 10000; i++ {
		for j := 2; j < 9; j++ {
			ok, value := isPanMultiple(i, j)
			if ok {
				a = append(a, value)
			}
		}
	}

	sort.Ints(a)
	fmt.Println(a)
}
