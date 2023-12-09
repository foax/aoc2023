package day01

import (
	"bufio"
	"io"
	"log/slog"
	"strings"
	"testing"
)

func TestPart1Handler(t *testing.T) {
	inputString := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	logger := slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{}))
	reader := strings.NewReader(inputString)
	input := readInput(logger, bufio.NewScanner(reader))
	want := 142
	got := part1Handler(logger, input)
	if got != want {
		t.Fatalf("part1Handler failed; wanted %d, got %d\n", want, got)
	}
}

func TestPart2Handler(t *testing.T) {
	inputString := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

	logger := slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{}))
	reader := strings.NewReader(inputString)
	input := readInput(logger, bufio.NewScanner(reader))
	want := 281
	got := part2Handler(logger, input)
	if got != want {
		t.Fatalf("part2Handler failed; wanted %d, got %d\n", want, got)
	}
}
