package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

	for i := 0; i < len(line); i++ {
		number := int(line[i] - '0')

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func main() {
	numbers, err := readMatrix("input.txt")

	if err != nil {
		panic(err)
	}

	total := make([]int, 0)
	carry := 0

	for i := len(numbers[0]) - 1; i >= 0; i-- {
		sum := carry
		for j := 0; j < len(numbers); j++ {
			sum += numbers[j][i]
		}

		total = append(total, sum%10)
		carry = sum / 10
	}

	for carry > 0 {
		total = append(total, carry%10)
		carry = carry / 10
	}

	for i := len(total) - 1; i >= 0; i-- {
		fmt.Print(total[i])
	}
}
