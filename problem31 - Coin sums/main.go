package main

import "fmt"

var coins []int

func do(a []int, index int, wanted int) int {
	sum := 0

	if index == 8 {
		if value(a) == wanted {
			return 1
		}
		return 0
	}

	limit := wanted / coins[index]
	for i := 0; i <= limit; i++ {
		clear(a, index+1)
		a[index] = i
		if value(a) > wanted {
			break
		}

		sum += do(a, index+1, wanted)
	}

	return sum
}

func clear(a []int, index int) {
	for i := index; i < len(a); i++ {
		a[i] = 0
	}
}

func value(a []int) int {
	sum := 0

	for i := 0; i < len(coins); i++ {
		sum += int(a[i]) * coins[i]
	}

	return sum
}

func main() {
	coins = []int{1, 2, 5, 10, 20, 50, 100, 200}
	wanted := 200
	a := make([]int, 8)

	fmt.Println(do(a, 0, wanted))
}
