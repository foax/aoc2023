package day09

import (
	"bufio"
	"log/slog"
	"strconv"
	"strings"
)

func readInput(scanner *bufio.Scanner) (output []string) {
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	slog.Debug("Lines of input read", "lines", len(output))
	return
}

func loadDataset(input []string) (dataset [][]int) {
	dataset = make([][]int, len(input))
	for i, line := range input {
		fields := strings.Fields(line)
		history := make([]int, len(fields))
		for j, field := range strings.Fields(line) {
			value, _ := strconv.Atoi(field)
			history[j] = value
		}
		dataset[i] = history
	}
	return
}

func nextSequence(x []int) (y []int) {
	y = make([]int, len(x)-1)
	for i := 0; i < len(x)-1; i++ {
		y[i] = x[i+1] - x[i]
	}
	return
}

func generateSequences(x []int) [][]int {
	sequences := make([][]int, 1)
	sequence := make([]int, len(x))
	copy(sequence, x)
	slog.Debug("generateSequences", "sequence", sequence, "x", x)
	sequences[0] = sequence
	for i, allZeros := 0, false; !allZeros; i++ {
		sequence = nextSequence(sequence)
		for j, y := range sequence {
			if y != 0 {
				break
			}
			if j == len(sequence)-1 {
				allZeros = true
				break
			}
		}
		sequences = append(sequences, sequence)
		slog.Debug("generateSequences", "i", i, "sequences", sequences)
	}
	return sequences
}

func nextHistoryValue(sequences [][]int) (nextValue int) {
	for i := len(sequences) - 1; i >= 0; i-- {
		nextValue += sequences[i][len(sequences[i])-1]
	}
	return
}

func previousHistoryValue(sequences [][]int) (prevValue int) {
	for i := len(sequences) - 1; i >= 0; i-- {
		prevValue = sequences[i][0] - prevValue
		slog.Debug("previousHistoryValue", "i", i, "prevValue", prevValue)
	}
	return
}

func part1Handler(input []string) (total int) {
	dataset := loadDataset(input)
	for i, d := range dataset {
		s := generateSequences(d)
		nextValue := nextHistoryValue(s)
		slog.Debug("part1Handler", "i", i, "d", d, "nextValue", nextValue)
		total += nextValue
	}
	return
}

func part2Handler(input []string) (total int) {
	dataset := loadDataset(input)
	for i, d := range dataset {
		s := generateSequences(d)
		prevValue := previousHistoryValue(s)
		slog.Debug("part2Handler", "i", i, "d", d, "prevValue", prevValue)
		total += prevValue
	}
	return
}

func Execute(scanner *bufio.Scanner) {
	input := readInput(scanner)
	part1Total := part1Handler(input)
	part2Total := part2Handler(input)
	slog.Info("Results", "part1", part1Total, "part2", part2Total)
}
