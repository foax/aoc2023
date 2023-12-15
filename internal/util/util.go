package util

import (
	"bufio"
	"log/slog"
)

func ReadInput(scanner *bufio.Scanner) (output []string) {
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	slog.Debug("Lines of input read", "lines", len(output))
	return
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func SliceSum(x []int) (sum int) {
	for _, y := range x {
		sum += y
	}
	return
}

func SumOfDescendingSequence(start, length int) int {
	// Formula to calculate sum of a descending sequence
	// Sum = n * (n + 1) / 2
	sum := (2*start + (length-1)*(-1)) * length / 2
	return sum
}

func RotateMatrix[T any](matrix [][]T, rotations int) [][]T {
	rows := len(matrix)
	cols := len(matrix[0])

	rotations = rotations % 4

	for r := 0; r < rotations; r++ {
		result := make([][]T, cols)
		for i := range result {
			result[i] = make([]T, rows)
		}

		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				result[j][rows-1-i] = matrix[i][j]
			}
		}

		matrix = result
		rows, cols = len(matrix), len(matrix[0])
	}

	return matrix
}

func TransposeMatrix[T any](matrix [][]T) [][]T {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]T, cols)
	for i := range transposed {
		transposed[i] = make([]T, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}
