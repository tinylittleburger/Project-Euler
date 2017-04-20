package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 998; i++ {
		for j := i; j <= 1000-i; j++ {
			k := 1000 - i - j

			if i*i+j*j == k*k {
				fmt.Printf("i:%v j:%v l:%v \n", i, j, k)
				fmt.Println(i * j * k)
				return
			}
		}
	}
}
