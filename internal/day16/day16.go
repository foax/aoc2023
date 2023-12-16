package day16

import (
	"bufio"
	"fmt"
	"log/slog"

	"github.com/foax/aoc2023/internal/util"
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
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

func (l LightContraption) AdvanceBeam(row, col int, direction int) {
	switch direction {
	case UP:
		row--
	case RIGHT:
		col++
	case DOWN:
		row++
	case LEFT:
		col--
	}
	l.BeamLight(row, col, direction)
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
		l.AdvanceBeam(row, col, direction)
	case '-':
		if direction == LEFT || direction == RIGHT {
			l.AdvanceBeam(row, col, direction)
		} else {
			l.AdvanceBeam(row, col, LEFT)
			l.AdvanceBeam(row, col, RIGHT)
		}
	case '|':
		if direction == UP || direction == DOWN {
			l.AdvanceBeam(row, col, direction)
		} else {
			l.AdvanceBeam(row, col, UP)
			l.AdvanceBeam(row, col, DOWN)
		}
	case '/':
		switch direction {
		case UP:
			l.AdvanceBeam(row, col, RIGHT)
		case RIGHT:
			l.AdvanceBeam(row, col, UP)
		case DOWN:
			l.AdvanceBeam(row, col, LEFT)
		case LEFT:
			l.AdvanceBeam(row, col, DOWN)
		}
	case '\\':
		switch direction {
		case UP:
			l.AdvanceBeam(row, col, LEFT)
		case RIGHT:
			l.AdvanceBeam(row, col, DOWN)
		case DOWN:
			l.AdvanceBeam(row, col, RIGHT)
		case LEFT:
			l.AdvanceBeam(row, col, UP)
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
				case t.beamed[UP]:
					x = "^"
				case t.beamed[RIGHT]:
					x = ">"
				case t.beamed[DOWN]:
					x = "v"
				case t.beamed[LEFT]:
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
	l.BeamLight(0, 0, RIGHT)
	return l.Energy()
}

func part2Handler(input []string) (total int) {
	for x := 0; x < len(input); x++ {
		for _, y := range [2][2]int{{0, RIGHT}, {len(input[0]) - 1, LEFT}} {
			l := Load(input)
			l.BeamLight(x, y[0], y[1])
			e := l.Energy()
			if e > total {
				total = e
			}
		}
	}
	for y := 0; y < len(input[0]); y++ {
		for _, x := range [2][2]int{{0, DOWN}, {len(input) - 1, UP}} {
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
