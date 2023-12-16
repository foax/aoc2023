package day16

import (
	"bufio"
	"fmt"
	"log/slog"

	"github.com/foax/aoc2023/internal/util"
)

type Tile struct {
	tileType rune
	beamed   [4]bool
}

type LightContraption [][]Tile

func Load(input []string) (l LightContraption) {
	l = make(LightContraption, len(input))
	for i := range input {
		l[i] = make([]Tile, len(input[i]))
		for j, x := range input[i] {
			l[i][j] = Tile{tileType: x}
		}
	}
	return l
}

func AdjustVector(row, col int, direction int) (int, int) {
	// 0 = up; 1 = right; 2 = down; 3 = left
	switch direction {
	case 0:
		return row - 1, col
	case 1:
		return row, col + 1
	case 2:
		return row + 1, col
	case 3:
		return row, col - 1
	}
	return 0, 0
}

func (l LightContraption) BeamLight(row, col int, direction int) {
	if row < 0 || row >= len(l) || col < 0 || col >= len(l[0]) {
		return
	}
	if l[row][col].beamed[direction] {
		return
	}
	l[row][col].beamed[direction] = true

	switch l[row][col].tileType {
	case '.':
		newRow, newCol := AdjustVector(row, col, direction)
		l.BeamLight(newRow, newCol, direction)
	case '-':
		if direction%2 == 1 {
			newRow, newCol := AdjustVector(row, col, direction)
			l.BeamLight(newRow, newCol, direction)
		} else {
			newRow, newCol := AdjustVector(row, col, 1)
			l.BeamLight(newRow, newCol, 1)
			newRow, newCol = AdjustVector(row, col, 3)
			l.BeamLight(newRow, newCol, 3)
		}
	case '|':
		if direction%2 == 0 {
			newRow, newCol := AdjustVector(row, col, direction)
			l.BeamLight(newRow, newCol, direction)
		} else {
			newRow, newCol := AdjustVector(row, col, 0)
			l.BeamLight(newRow, newCol, 0)
			newRow, newCol = AdjustVector(row, col, 2)
			l.BeamLight(newRow, newCol, 2)
		}
	case '/':
		switch direction {
		case 0:
			newRow, newCol := AdjustVector(row, col, 1)
			l.BeamLight(newRow, newCol, 1)
		case 1:
			newRow, newCol := AdjustVector(row, col, 0)
			l.BeamLight(newRow, newCol, 0)
		case 2:
			newRow, newCol := AdjustVector(row, col, 3)
			l.BeamLight(newRow, newCol, 3)
		case 3:
			newRow, newCol := AdjustVector(row, col, 2)
			l.BeamLight(newRow, newCol, 2)
		}
	case '\\':
		switch direction {
		case 0:
			newRow, newCol := AdjustVector(row, col, 3)
			l.BeamLight(newRow, newCol, 3)
		case 1:
			newRow, newCol := AdjustVector(row, col, 2)
			l.BeamLight(newRow, newCol, 2)
		case 2:
			newRow, newCol := AdjustVector(row, col, 1)
			l.BeamLight(newRow, newCol, 1)
		case 3:
			newRow, newCol := AdjustVector(row, col, 0)
			l.BeamLight(newRow, newCol, 0)
		}
	}
}

func PrintContraption(l LightContraption) {
	for _, line := range l {
		lineStr := ""
		for _, t := range line {
			totalBeamed := 0
			var x string
			for _, b := range t.beamed {
				if b {
					totalBeamed++
				}
			}
			switch {
			case t.tileType == '/' || t.tileType == '\\' || t.tileType == '|' || t.tileType == '-':
				x = string(t.tileType)
			case totalBeamed == 1:
				switch {
				case t.beamed[0]:
					x = "^"
				case t.beamed[1]:
					x = ">"
				case t.beamed[2]:
					x = "v"
				case t.beamed[3]:
					x = "<"
				}
			case totalBeamed > 1:
				x = fmt.Sprintf("%d", totalBeamed)
			default:
				x = "."
			}
			lineStr += x
		}
		fmt.Printf("%s\n", lineStr)
	}
}

func (l LightContraption) Energy() (total int) {
	for _, line := range l {
		for _, t := range line {
			totalBeamed := 0
			for _, b := range t.beamed {
				if b {
					totalBeamed++
				}
			}
			if totalBeamed > 0 {
				total++
			}
		}
	}
	return
}

func part1Handler(input []string) int {
	l := Load(input)
	l.BeamLight(0, 0, 1)
	return l.Energy()
}

func part2Handler(input []string) (total int) {
	// bare := Load(input)
	for x := 0; x < len(input); x++ {
		for _, y := range [2][2]int{{0, 1}, {len(input[0]) - 1, 3}} {
			l := Load(input)
			l.BeamLight(x, y[0], y[1])
			e := l.Energy()
			if e > total {
				total = e
			}
		}
	}
	for y := 0; y < len(input[0]); y++ {
		for _, x := range [2][2]int{{0, 2}, {len(input) - 1, 0}} {
			l := Load(input)
			l.BeamLight(x[0], y, x[1])
			e := l.Energy()
			if e > total {
				total = e
			}
		}
	}
	return total
}

func Execute(scanner *bufio.Scanner) {
	input := util.ReadInput(scanner)
	part1 := part1Handler(input)
	part2 := part2Handler(input)
	slog.Info("Results", "part1", part1, "part2", part2)
}
