package main

import "sort"
import "fmt"

func nextPermutation(a []int) []int {
	for right := len(a) - 1; right > 0; right-- {
		left := right - 1

		if a[right] > a[left] {
			indexSmallest := right + findSmallestGreater(a[left], a[right:])
			swap(left, indexSmallest, a)
			subA := a[right:]
			sort.Ints(subA)
			return a
		}
	}

	return nil
}

func findSmallestGreater(value int, a []int) int {
	smallest := 0

	for i, v := range a {
		if v > value && v < a[smallest] {
			smallest = i
		}
	}

	return smallest
}

func swap(i int, j int, a []int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func isSubPan(a []int) bool {
	primes := []int{17, 13, 11, 7, 5, 3, 2}

	for i, value := range primes {
		if !isDivisible(a, len(a)-3-i, value) {
			return false
		}
	}

	return true
}

func isDivisible(a []int, index int, d int) bool {
	n := a[index]*100 + a[index+1]*10 + a[index+2]

	return n%d == 0
}

func toNumber(a []int) int64 {
	n := int64(0)

	for _, value := range a {
		n = n*10 + int64(value)
	}

	return n
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.Ints(a)
	sum := int64(0)

	for ; a != nil; a = nextPermutation(a) {
		if isSubPan(a) {
			fmt.Println(a)
			sum += toNumber(a)
		}
	}

	fmt.Println(sum)
}
