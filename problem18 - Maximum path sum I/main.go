package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readMatrix(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("File opening failed")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix := make([][]int, 0)

	for scanner.Scan() {
		line, err := readArray(scanner.Text())
		if err != nil {
			return nil, errors.New("Reading line failed")
		}

		matrix = append(matrix, line)
	}

	return matrix, nil
}

func readArray(line string) ([]int, error) {
	numbers := make([]int, 0)

	for _, value := range strings.Split(line, " ") {
		number, err := strconv.Atoi(value)

		if err != nil {
			return nil, errors.New("Converting to int failed")
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func maxRoute(m [][]int, n [][]int, d int, p int) {
	if d == len(m)-1 {
		n[d][p] = m[d][p]
		return
	}

	if n[d+1][p] == 0 {
		maxRoute(m, n, d+1, p)
	}

	if n[d+1][p+1] == 0 {
		maxRoute(m, n, d+1, p+1)
	}

	n[d][p] = m[d][p] + max(n[d+1][p], n[d+1][p+1])
}

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func main() {
	m, err := readMatrix("input.txt")

	if err != nil {
		panic(err)
	}

	n := make([][]int, len(m))
	for i := 0; i < len(n); i++ {
		n[i] = make([]int, i+1)
	}

	maxRoute(m, n, 0, 0)
	fmt.Println(n[0][0])
}
