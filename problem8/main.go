package main

import "fmt"
import "io/ioutil"

func main() {
	frame := 13

	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	number := string(b)

	var max int64
	size := len(number)

	for i := 0; i < size-frame; i++ {
		var product int64 = 1

		for j := 0; j < frame; j++ {
			product *= int64(number[i+j] - '0')
		}

		if int64(product) > max {
			max = int64(product)
		}
	}

	fmt.Println(max)
}
