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

func loadGalaxy(input []string, f int) galaxy {
	g := galaxy{emptyFactor: f}
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
	return g
}

func galaxySteps(g galaxy, x, y [2]int) int {

	steps := util.IntAbs(x[0]-y[0]) + util.IntAbs(x[1]-y[1])
	for _, data := range []struct {
		idx   int
		empty map[int]bool
	}{
		{0, g.emptyRow},
		{1, g.emptyCol},
	} {

		step := 0
		if x[data.idx] > y[data.idx] {
			step = -1
		} else if x[data.idx] < y[data.idx] {
			step = 1
		} else {
			step = 0
		}

		for a := x[data.idx] + step; a != y[data.idx]; a += step {
			if data.empty[a] {
				steps += g.emptyFactor - 1
			}
		}
	}
	return steps

}

func partHandler(input []string, f int) int {
	g := loadGalaxy(input, f)
	total := 0
	for i, x := range g.coords {
		for _, y := range g.coords[i+1:] {
			steps := galaxySteps(g, x, y)
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
