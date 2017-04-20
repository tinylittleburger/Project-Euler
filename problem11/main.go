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

func main() {
	frame := 2
	matrix, err := readMatrix("input.txt")

	if err != nil {
		panic(err)
	}

	mRows := len(matrix)
	mColumns := len(matrix[0])

	for i := 0; i < mRows; i++ {
		if len(matrix[i]) != mColumns {
			fmt.Println("Not a proper matrix, exiting...")
			return
		}
	}

	if frame > mRows || frame > mColumns {
		fmt.Println("Frame too big, exiting...")
		return
	}

	var max int64
	for i := 0; i <= mColumns-frame; i++ {

		for j := 0; j <= mColumns-frame; j++ {
			value := matrix[i][j]
			diag := int64(value)
			diag2 := int64(matrix[i][j+frame-1])
			horiz := int64(value)
			vert := int64(value)

			for k := 1; k <= frame-1; k++ {
				horiz *= int64(matrix[i][j+k])
				vert *= int64(matrix[i+k][j])
				diag *= int64(matrix[i+k][j+k])
				diag2 *= int64(matrix[i+k][j+frame-1-k])
			}

			if diag > max {
				max = diag
			}

			if diag2 > max {
				max = diag2
			}

			if horiz > max {
				max = horiz
			}

			if vert > max {
				max = vert
			}
		}
	}

	for i := mRows - frame + 1; i < mRows; i++ {
		for j := 0; j <= mColumns-frame; j++ {
			value := matrix[i][j]
			horiz := int64(value)

			for k := 1; k <= frame-1; k++ {
				horiz *= int64(matrix[i][j+k])
			}

			if horiz > max {
				max = horiz
			}
		}
	}

	for i := 0; i <= mColumns-frame; i++ {
		for j := mColumns - frame + 1; j < mColumns; j++ {
			value := matrix[i][j]
			vert := int64(value)

			for k := 1; k <= frame-1; k++ {
				vert *= int64(matrix[i+k][j])
			}

			if vert > max {
				max = vert
			}
		}
	}

	fmt.Println(max)
}
