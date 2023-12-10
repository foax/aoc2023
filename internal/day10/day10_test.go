package day10

import (
	"bufio"
	"strings"
	"testing"

	"github.com/foax/aoc2023/internal/util"
)

var inputString string = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

func TestFindStart(t *testing.T) {
	reader := strings.NewReader(inputString)
	input := util.ReadInput(bufio.NewScanner(reader))
	grid := initGrid(len(input), len(input[0]))
	for i, line := range input {
		for j, pipe := range line {
			grid[i][j] = pipe
		}
	}
	want := [2]int{2, 0}
	got := findStart(grid)
	if want != got {
		t.Errorf("findStart failed: wanted %v, got %v", want, got)
	}
}

func TestLoopLength(t *testing.T) {
	input := []string{"7-F7-", ".FJ|7", "SJLL7", "|F--J", "LJ.LJ"}
	grid := initGrid(len(input), len(input[0]))
	for i, line := range input {
		for j, pipe := range line {
			grid[i][j] = pipe
		}
	}
	want := 16
	got, _ := loopLength(grid, [2]int{2, 0})
	if want != got {
		t.Errorf("findStart failed: wanted %v, got %v", want, got)
	}
}
