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
