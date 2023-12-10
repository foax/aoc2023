package day01

import (
	"bufio"
	"log/slog"
	"regexp"
	"strings"
)

func part1Handler(input []string) (total int) {
	for _, line := range input {
		slog.Debug("part1Handler", "line", line)
		calibration := make([]int, 0)
		for _, c := range line {
			if x := int(c) - '0'; x >= 0 && x <= 9 {
				calibration = append(calibration, x)
			}
		}
		slog.Debug("part1Handler", "calibration", calibration)
		total += 10*calibration[0] + calibration[len(calibration)-1]
	}
	return
}

func part2Handler(input []string) (total int) {
	var numberMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	numberMapKeys := make([]string, 0)
	for k := range numberMap {
		numberMapKeys = append(numberMapKeys, k)
	}
	re, _ := regexp.Compile(strings.Join(numberMapKeys, "|"))
	slog.Debug("part2Handler", "re", re)

	for _, line := range input {
		slog.Debug("part2Handler", "line", line)
		numberSet := make([]bool, len(line))
		numberList := make([]int, len(line))

		// find worded numbers in the line
		for i := 0; i < len(line); {
			n := re.FindStringIndex(line[i:])
			slog.Debug("part2Handler", "n", n)
			if n == nil {
				break
			}
			numberSet[n[0]+i] = true
			numberList[n[0]+i] = numberMap[line[n[0]+i:n[1]+i]]
			i += n[0] + 1
		}
		slog.Debug("part2Handler: word parsing done", "numberSet", numberSet, "numberList", numberList)

		// find numerical runes in the line
		for i, c := range line {
			x := int(c) - '0'
			if x >= 0 && x <= 9 {
				numberSet[i] = true
				numberList[i] = x
			}
		}

		slog.Debug("part2Handler: digit parsing done", "numberSet", numberSet, "numberList", numberList)

		calibration := make([]int, 0)
		for i, c := range numberSet {
			if c {
				calibration = append(calibration, numberList[i])
			}
		}

		total += 10*calibration[0] + calibration[len(calibration)-1]
		slog.Debug("part2Handler", "calibration", calibration, "total", total)
	}
	slog.Debug("part2Handler", "total", total)
	return
}

func readInput(scanner *bufio.Scanner) (output []string) {
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	slog.Debug("Lines of input read", "lines", len(output))
	return
}

func Execute(scanner *bufio.Scanner) {
	input := readInput(scanner)
	part1Total := part1Handler(input)
	part2Total := part2Handler(input)
	slog.Info("Part 1", "result", part1Total)
	slog.Info("Part 2", "result", part2Total)
}
