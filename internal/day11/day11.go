package day11

import (
	"bufio"
	"log/slog"

	"github.com/foax/aoc2023/internal/util"
)

type galaxy struct {
	coords      [][2]int
	emptyRow    map[int]bool
	emptyCol    map[int]bool
	emptyFactor int
}

func (g *galaxy) Load(input []string, f int) {
	g.emptyFactor = f
	g.emptyRow = make(map[int]bool)
	g.emptyCol = make(map[int]bool)

	galRows := make(map[int]bool)
	galCols := make(map[int]bool)
	for i, line := range input {
		for j, star := range line {
			if star == '#' {
				g.coords = append(g.coords, [2]int{i, j})
				galRows[i] = true
				galCols[j] = true
			}
		}
	}

	for i := 0; i < len(input); i++ {
		if !galRows[i] {
			g.emptyRow[i] = true
		}
	}
	for j := 0; j < len(input[0]); j++ {
		if !galCols[j] {
			g.emptyCol[j] = true
		}
	}
}

func (g *galaxy) Steps(a, b [2]int) int {
	steps := util.IntAbs(a[0]-b[0]) + util.IntAbs(a[1]-b[1])
	for _, data := range []struct {
		idx   int
		empty map[int]bool
	}{
		{0, g.emptyRow},
		{1, g.emptyCol},
	} {

		step := 0
		if a[data.idx] > b[data.idx] {
			step = -1
		} else if a[data.idx] < b[data.idx] {
			step = 1
		} else {
			step = 0
		}

		for x := a[data.idx] + step; x != b[data.idx]; x += step {
			if data.empty[x] {
				steps += g.emptyFactor - 1
			}
		}
	}
	return steps
}

func partHandler(input []string, f int) int {
	g := galaxy{}
	g.Load(input, f)
	total := 0
	for i, a := range g.coords {
		for _, b := range g.coords[i+1:] {
			steps := g.Steps(a, b)
			total += steps
		}
	}
	return total
}

func Execute(scanner *bufio.Scanner) {
	input := util.ReadInput(scanner)
	part1 := partHandler(input, 2)
	part2 := partHandler(input, 1000000)

	slog.Info("Results", "part1", part1, "part2", part2)
}
