package main

import (
	"fmt"
)

func divSum(n int) (sum int) {
	sum = 1
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

	return
}

func main() {
	limit := 10000
	a := make([]int, limit)
	amicable := 0

	for i := 1; i < limit; i++ {
		a[i] = divSum(i)

		first := a[i]
		if first < limit {
			second := a[a[i]]

			if first != i && second != 0 && second == i {
				amicable += a[i] + i
			}
		}
	}

	fmt.Println(amicable)
}
