package main

import (
	"fmt"
)

func main() {
	var number int64 = 600851475143
	var factor int64 = 2
	var bpf int64 = 1

	for factor <= number {
		if number%factor == 0 {
			number /= factor
			bpf = factor
		} else {
			factor++
		}
	}

	fmt.Println(bpf)
}
