package day10

import (
	"bufio"
	"fmt"
	"log/slog"

	"github.com/foax/aoc2023/internal/util"
)

func findStart(grid []string) [2]int {
	for i, line := range grid {
		for j, pipe := range line {
			if pipe == 'S' {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

func initGrid(x int, y int) [][]rune {
	grid := make([][]rune, x)
	for i := range grid {
		grid[i] = make([]rune, y)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}

func loopLength(grid []string, start [2]int) (int, [][]rune) {
	slog.Debug("loopLength", "grid", grid)
	var oldLoc, pipeLoc [2]int
	var startingPipes [4][3]rune = [4][3]rune{
		{'|', '7', 'F'},
		{'-', '7', 'J'},
		{'|', 'J', 'L'},
		{'-', 'F', 'L'},
	}
	loopGrid := initGrid(len(grid), len(grid[0]))
	loopGrid[start[0]][start[1]] = 'S'
	length := 1
Loop:
	for i, delta := range [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		pipeLoc = [2]int{start[0] + delta[0], start[1] + delta[1]}
		slog.Debug("loopLength", "i", i, "delta", delta, "pipeLoc", pipeLoc)
		if pipeLoc[0] < 0 || pipeLoc[0] >= len(grid) || pipeLoc[1] < 0 || pipeLoc[1] >= len(grid[pipeLoc[0]]) {
			continue
		}
		for _, p := range startingPipes[i] {
			slog.Debug("loopLength", "p", string(p))
			if p == rune(grid[pipeLoc[0]][pipeLoc[1]]) {
				oldLoc = start
				break Loop
			}
		}
	}
	loopGrid[pipeLoc[0]][pipeLoc[1]] = rune(grid[pipeLoc[0]][pipeLoc[1]])

	for pipeLoc != start {
		tmpLoc := pipeLoc
		switch rune(grid[pipeLoc[0]][pipeLoc[1]]) {
		case '|':
			if oldLoc[0] < pipeLoc[0] {
				pipeLoc[0]++
			} else {
				pipeLoc[0]--
			}
		case '-':
			if oldLoc[1] < pipeLoc[1] {
				pipeLoc[1]++
			} else {
				pipeLoc[1]--
			}
		case 'L':
			if pipeLoc[0] == oldLoc[0] {
				pipeLoc[0]--
			} else {
				pipeLoc[1]++
			}
		case 'J':
			if pipeLoc[0] == oldLoc[0] {
				pipeLoc[0]--
			} else {
				pipeLoc[1]--
			}
		case 'F':
			if pipeLoc[0] == oldLoc[0] {
				pipeLoc[0]++
			} else {
				pipeLoc[1]++
			}
		case '7':
			if pipeLoc[0] == oldLoc[0] {
				pipeLoc[0]++
			} else {
				pipeLoc[1]--
			}
		}
		oldLoc = tmpLoc
		loopGrid[pipeLoc[0]][pipeLoc[1]] = rune(grid[pipeLoc[0]][pipeLoc[1]])
		length++
	}
	return length, loopGrid
}

func gridNeighbour(grid [][]rune, pos [2]int, delta [2]int) rune {
	var check [2]int = [2]int{pos[0] + delta[0], pos[1] + delta[1]}
	if check[0] < 0 || check[0] >= len(grid) || check[1] < 0 || check[1] >= len(grid[0]) {
		return 'O'
	}
	return grid[check[0]][check[1]]
}

// expand a grid so it's easier to detect areas enclosed by the loop
func expandGrid(grid [][]rune) [][]rune {
	expandedGrid := initGrid(len(grid)*2-1, len(grid[0])*2-1)

	// start by expanding each line; add a '-' if it's joining two pipes
	for i, line := range grid {
		for j, pipe := range line {
			expandedGrid[i*2][j*2] = pipe
			if j == len(line)-1 {
				break
			}
			right := gridNeighbour(grid, [2]int{i, j}, [2]int{0, 1})
			if (pipe == '-' || pipe == 'L' || pipe == 'F' || pipe == 'S') && (right == '-' || right == '7' || right == 'J' || right == 'S') {
				expandedGrid[i*2][j*2+1] = '-'
			}
		}
	}

	for i, line := range expandedGrid {
		if i%2 == 0 {
			continue
		}
		for j := range line {
			above := gridNeighbour(expandedGrid, [2]int{i, j}, [2]int{-1, 0})
			below := gridNeighbour(expandedGrid, [2]int{i, j}, [2]int{1, 0})
			if (above == '|' || above == 'F' || above == '7' || above == 'S') && (below == '|' || below == 'L' || below == 'J' || below == 'S') {
				expandedGrid[i][j] = '|'
			}
		}
	}

	//
	return expandedGrid
}

func findOutsidePipes(grid [][]rune) {
	for i, line := range grid {
		for j := range line {
			if grid[i][j] != '.' {
				continue
			}
			for _, delta := range [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
				y := gridNeighbour(grid, [2]int{i, j}, delta)
				if y == 'O' {
					markOutsidePipe(grid, [2]int{i, j})
				}
			}
		}
	}
}

func markOutsidePipe(grid [][]rune, pos [2]int) {
	grid[pos[0]][pos[1]] = 'O'
	for _, delta := range [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		y := gridNeighbour(grid, pos, delta)
		if y == '.' {
			// recursion...
			markOutsidePipe(grid, [2]int{pos[0] + delta[0], pos[1] + delta[1]})
		}
	}
}

func countInsidePipes(grid [][]rune) (total int) {
	for i, line := range grid {
		for j, pipe := range line {
			if pipe != '.' {
				continue
			}
			fourDots := true
			for _, delta := range [4][2]int{{0, 1}, {1, 1}, {1, 0}} {
				y := gridNeighbour(grid, [2]int{i, j}, delta)
				if y != '.' {
					fourDots = false
					break
				}
			}
			if fourDots {
				for _, delta := range [4][2]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}} {
					grid[i+delta[0]][j+delta[1]] = 'I'
				}
				total++
			}
		}
	}
	return total
}

func part1Handler(grid []string) int {
	length, _ := loopLength(grid, findStart(grid))
	return length / 2
}

func part2Handler(grid []string) int {
	_, loopGrid := loopLength(grid, findStart(grid))
	for _, line := range loopGrid {
		fmt.Println(string(line))
	}
	fmt.Println()
	expandedGrid := expandGrid(loopGrid)
	for _, line := range expandedGrid {
		fmt.Println(string(line))
	}

	findOutsidePipes(expandedGrid)
	for _, line := range expandedGrid {
		fmt.Println(string(line))
	}

	total := countInsidePipes(expandedGrid)
	for _, line := range expandedGrid {
		fmt.Println(string(line))
	}
	return total
}

func Execute(scanner *bufio.Scanner) {
	input := util.ReadInput(scanner)
	// part1Total := part1Handler(input)
	part2Total := part2Handler(input)
	slog.Info("Results", "part2", part2Total)
}
