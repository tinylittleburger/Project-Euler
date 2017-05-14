package utilities

import "sort"

func NextPermutation(a []int) []int {
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

func ToSlice(n int) []int {
	a := []int{}

	for n > 0 {
		a = append(a, n%10)
		n /= 10
	}

	for i := 0; i*2 < len(a); i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}

	return a
}

func ToNumber(a []int) int {
	n := 0

	for _, v := range a {
		n = n*10 + v
	}

	return n
}
