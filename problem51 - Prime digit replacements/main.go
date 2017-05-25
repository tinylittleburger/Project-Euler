package main

import (
	"comi/Project-Euler/utilities"
	"fmt"
	"strconv"
	"strings"
)

func countPrimes(s string, sieve *utilities.Sieve) (int, int) {
	var smallest, count int
	length := len(s)

	for i := 0; i <= 9; i++ {
		s := s

		for j := 0; j < length; j++ {
			if s[j] == '*' && !(i == 0 && j == 0) {
				s = replaceAtIndex(s, strconv.Itoa(i), j)
			}
		}

		n, _ := strconv.Atoi(s)

		if sieve.IsPrime(n) {
			if count == 0 {
				smallest = n
			}
			count++
		}
	}

	return count, smallest
}

func createMasks(s string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			return createMasksInternal(s, i)
		}
	}

	return nil
}

func replaceAtIndex(s string, l string, index int) string {
	return s[:index] + l + s[index+1:]
}

func createMasksInternal(s string, index int) []string {
	var a []string
	next := -1

	for i := index + 1; i < len(s); i++ {
		if s[i] == '0' {
			next = i
			break
		}
	}

	for j := 0; j <= 9; j++ {
		s := replaceAtIndex(s, strconv.Itoa(j), index)

		if next == -1 {
			a = append(a, s)
		} else {
			b := createMasksInternal(s, next)
			a = append(a, b...)
		}
	}

	return a
}

func replaceOnes(s string) string {
	return strings.Replace(s, "1", "*", -1)
}

func main() {
	sieve := utilities.NewSieve()
	sieve.Init(1000000)

	for i := 2; ; i = i + 2 {
		s := strconv.FormatInt(int64(i), 2)
		s = replaceOnes(s)

		masks := createMasks(s)

		for _, mask := range masks {
			primes, smallest := countPrimes(mask, &sieve)
			if primes == 8 {
				fmt.Println(smallest)
				return
			}
		}
	}
}
