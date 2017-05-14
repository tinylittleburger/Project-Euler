package main

import "math"
import "fmt"

func findDigitInternal(pos, digs int) int {
	lowBound := int(math.Pow10(digs - 1))
	r := int(math.Pow10(digs)) - lowBound

	if pos > digs*r {
		return findDigitInternal(pos-digs*r, digs+1)
	}

	number := lowBound + pos/digs - 1
	var divs int
	var digit int

	if pos%digs == 0 {
		divs = 0
	} else {
		number++
		divs = digs - pos%digs
	}

	for i := 0; i <= divs; i++ {
		digit = number % 10
		number /= 10
	}

	return digit
}

func findDigit(pos int) int {
	return findDigitInternal(pos, 1)
}

func main() {
	a := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	p := 1

	for _, value := range a {
		p *= findDigit(value)
	}

	fmt.Println(p)
}
