package day01

import (
	"bufio"
	"io"
	"log/slog"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	logger := slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{}))
	slog.SetDefault(logger)
	m.Run()
}

func TestPart1Handler(t *testing.T) {
	inputString := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	reader := strings.NewReader(inputString)
	input := readInput(bufio.NewScanner(reader))
	want := 142
	got := part1Handler(input)
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

	reader := strings.NewReader(inputString)
	input := readInput(bufio.NewScanner(reader))
	want := 281
	got := part2Handler(input)
	if got != want {
		t.Fatalf("part2Handler failed; wanted %d, got %d\n", want, got)
	}
}
