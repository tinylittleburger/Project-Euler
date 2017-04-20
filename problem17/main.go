package main

import "fmt"

func spellNumber(n int) string {
	basic := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight",
		"nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen",
		"sixteen", "seventeen", "eighteen", "nineteen"}

	tens := []string{"", "", "twenty", "thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninety"}

	pows := []string{"", "", "hundred", "thousand"}

	var ns string

	if n >= 1000 {
		digit := n / 1000
		ns += basic[digit] + pows[3]
		n %= 1000
	}

	if n >= 100 {
		digit := n / 100
		ns += basic[digit] + pows[2]
		n %= 100

		if n != 0 {
			ns += "and"
		}
	}

	if n >= 20 {
		digit := n / 10
		ns += tens[digit]
		n %= 10
	}

	if n > 0 {
		ns += basic[n]
	}

	return ns
}

func main() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += len(spellNumber(i))
	}

	fmt.Println(sum)
}
