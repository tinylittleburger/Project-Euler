package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func triangle(n int) int {
	return n * (n + 1) / 2
}

func alphIndex(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	}

	if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 1
	}

	return 0
}

func wordValue(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += alphIndex(s[i])
	}

	return sum
}

func read(fileName string) ([]string, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("File opening failed: %v", err)
	}

	names := []string{}
	for _, value := range strings.Split(string(file), ",") {
		names = append(names, strings.Trim(value, "\""))
	}

	return names, nil
}

func main() {
	a, err := read("input.txt")

	if err != nil {
		panic(err)
	}

	m := make(map[int]int)
	sum := 0

	for i := 1; i < 50; i++ {
		m[triangle(i)]++
	}

	for _, w := range a {
		if _, ok := m[wordValue(w)]; ok {
			sum++
		}
	}

	fmt.Println(sum)
}
