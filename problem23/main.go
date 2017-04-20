package main

import (
	"fmt"
)

func sieve(a []bool) {
	for i := 2; i < len(a); i++ {
		if !a[i] {
			for j := 2; i*j < len(a); j++ {
				a[i*j] = true
			}
		}
	}
}

func isAbundant(n int, abundant []int, notprime []bool) bool {
	if abundant[n] != 0 {
		return abundant[n] > 0
	}

	if !notprime[n] {
		abundant[n] = -1
		return false
	}

	if divSum(n) > n {
		abundant[n] = 1
		return true
	}

	abundant[n] = -1
	return false
}

func divSum(n int) int {
	sum := 1
	i := 2
	for ; i*i < n; i++ {
		if n%i == 0 {
			sum += i
			sum += n / i
		}
	}

	if i*i == n {
		sum += i
	}

	return sum
}

func main() {
	size := 28123
	notprime := make([]bool, size+1)
	abundant := make([]int, size+1)
	sieve(notprime)

	sum := int64(0)

	for i := 1; i <= size; i++ {
		ab := false
		for j := 1; j*2 <= i; j++ {
			first := j
			second := i - j

			if isAbundant(first, abundant, notprime) && isAbundant(second, abundant, notprime) {
				ab = true
			}
		}
		if !ab {
			sum += int64(i)
		}
	}

	fmt.Println(sum)
}
