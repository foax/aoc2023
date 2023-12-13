package day13

import (
	"bufio"
	"log/slog"
	"reflect"

	"github.com/foax/aoc2023/internal/util"
)

type lavaMap [][]bool

func (l lavaMap) String() (s string) {
	for _, x := range l {
		for _, y := range x {
			switch y {
			case true:
				s += "#"
			case false:
				s += "."
			}
		}
		s += "\n"
	}
	return
}

func readLavaMaps(input []string) (maps []lavaMap) {
	lm := make(lavaMap, 0)
	for _, line := range input {
		if len(line) == 0 {
			maps = append(maps, lm)
			lm = make(lavaMap, 0)
			continue
		}
		l := make([]bool, len(line))
		for i, x := range line {
			l[i] = x == '#'
		}
		lm = append(lm, l)
	}
	maps = append(maps, lm)
	return
}

func rotateLavaMap(lm lavaMap) (rotated lavaMap) {
	rotated = make(lavaMap, len(lm[0]))
	for x := 0; x < len(lm[0]); x++ {
		rotated[x] = make([]bool, len(lm))
		for y := 0; y < len(lm); y++ {
			rotated[x][y] = lm[y][x]
		}
	}
	return
}

func checkReflection(lm lavaMap, pivot int) bool {
	lenLm := len(lm)
	for x, y := pivot, pivot+1; x >= 0 && y < lenLm; x, y = x-1, y+1 {
		if !reflect.DeepEqual(lm[x], lm[y]) {
			return false
		}
	}
	return true
}

func checkAllReflections(lm lavaMap) int {
	for i := 0; i < len(lm)-1; i++ {
		if checkReflection(lm, i) {
			return i + 1
		}
	}
	return 0
}

func part1Handler(input []string) (total int) {
	lavaMaps := readLavaMaps(input)
	for _, l := range lavaMaps {
		x := checkAllReflections(l)
		if x > 0 {
			total += 100 * x
			continue
		}
		x = checkAllReflections(rotateLavaMap(l))
		total += x
	}
	return total
}

func Execute(scanner *bufio.Scanner) {
	input := util.ReadInput(scanner)
	part1 := part1Handler(input)
	slog.Info("Results", "part1", part1)
}
