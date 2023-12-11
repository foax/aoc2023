package day11

import (
	"bufio"
	"fmt"

	"github.com/foax/aoc2023/internal/util"
)

type galaxy [][]rune

func printGalaxy(g galaxy) {
	for _, line := range g {
		fmt.Println(string(line))
	}
}

func loadGalaxy(input []string) galaxy {
	g := make(galaxy, len(input))
	for i, line := range input {
		g[i] = []rune(line)
	}
	return g
}

func expandGalaxy(g galaxy) (galaxy, [][2]int) {
	galRows := make(map[int]bool)
	galCols := make(map[int]bool)

	// check how big this new galaxy needs to be

	for i, line := range g {
		for j, x := range line {
			if x == '#' {
				galRows[i] = true
				galCols[j] = true
			}
		}
	}

	fmt.Println(galRows, galCols)

	newGalaxy := make(galaxy, 0)
	galaxyCoords := make([][2]int, 0)
	lineOfDots := make([]rune, 2*len(g[0])-len(galCols))
	for i := range lineOfDots {
		lineOfDots[i] = '.'
	}

	for i, line := range g {
		if !galRows[i] {
			newGalaxy = append(newGalaxy, lineOfDots, lineOfDots)
		} else {
			newLine := make([]rune, 2*len(g[0])-len(galCols))
			for j, k := 0, 0; k < len(newLine); j++ {
				if !galCols[j] {
					newLine[k] = '.'
					k++
					newLine[k] = '.'
				} else {
					newLine[k] = line[j]
					if newLine[k] == '#' {
						galaxyCoords = append(galaxyCoords, [2]int{len(newGalaxy), k})
					}
				}
				k++
			}
			newGalaxy = append(newGalaxy, newLine)
		}
	}
	return newGalaxy, galaxyCoords
}

func galaxySteps(x, y [2]int) int {
	return util.IntAbs(x[0]-y[0]) + util.IntAbs(x[1]-y[1])
}

func Execute(scanner *bufio.Scanner) {
	input := util.ReadInput(scanner)
	g := loadGalaxy(input)
	printGalaxy(g)
	fmt.Println("---")
	g, coords := expandGalaxy(g)
	printGalaxy(g)
	fmt.Println(coords)

	total := 0
	for i, x := range coords {
		for _, y := range coords[i+1:] {
			steps := galaxySteps(x, y)
			fmt.Printf("Steps from %v to %v: %d\n", x, y, steps)
			total += steps
		}
	}
	fmt.Println(total)

	// slog.Info("Results", "part1", part1Total, "part2", part2Total)
}
