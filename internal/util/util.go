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
