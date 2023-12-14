package day13

import (
	"bufio"
	"fmt"
	"log/slog"

	"github.com/foax/aoc2023/internal/util"
)

type lavaMap []string
type results struct {
	rotated bool
	pivot   int
}

func (l lavaMap) String() (s string) {
	for _, line := range l {
		s += fmt.Sprintln(line)
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
		lm = append(lm, line)
	}
	maps = append(maps, lm)
	return
}

func (lm lavaMap) Rotate() (rotated lavaMap) {
	rotated = make(lavaMap, len(lm[0]))
	for x := 0; x < len(lm[0]); x++ {
		rotated[x] = ""
		for y := 0; y < len(lm); y++ {
			rotated[x] += string(lm[y][x])
		}
	}
	return
}

// return how many elements are different
func compareLines(a string, b string) int {
	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}

func (lm lavaMap) CheckReflection(pivot int, smudge bool) bool {
	lenLm := len(lm)
	slog.Debug("checkReflection", "pivot", pivot, "smudge", smudge, "len", lenLm)
	for x, y := pivot, pivot+1; x >= 0 && y < lenLm; x, y = x-1, y+1 {
		count := compareLines(lm[x], lm[y])
		slog.Debug("checkReflection", "x", x, "y", y, "count", count)
		if (smudge && count > 1) || (!smudge && count > 0) {
			return false

		}
	}
	return true
}

func (lm lavaMap) FindReflection(smudge bool, skip int) int {
	for i := 0; i < len(lm)-1; i++ {
		if i == skip {
			continue
		}
		if lm.CheckReflection(i, smudge) {
			return i + 1
		}
	}
	return 0
}

func part1Handler(input []string) (total int, r []results) {
	lavaMaps := readLavaMaps(input)
	r = make([]results, 0)
	for _, l := range lavaMaps {
		x := l.FindReflection(false, -1)
		if x > 0 {
			total += 100 * x
			r = append(r, results{false, x})
			continue
		}
		x = l.Rotate().FindReflection(false, -1)
		total += x
		r = append(r, results{true, x})
	}
	return
}

func part2Handler(input []string, r []results) (total int) {
	var skip int
	lavaMaps := readLavaMaps(input)

	for i, l := range lavaMaps {
		if !r[i].rotated {
			skip = r[i].pivot - 1
		} else {
			skip = -1
		}
		x := l.FindReflection(true, skip)
		if x > 0 {
			total += 100 * x
			continue
		}
		if r[i].rotated {
			skip = r[i].pivot - 1
		} else {
			skip = -1
		}
		x = l.Rotate().FindReflection(true, skip)
		total += x
	}
	return total
}

func Execute(scanner *bufio.Scanner) {
	input := util.ReadInput(scanner)
	part1, r := part1Handler(input)
	part2 := part2Handler(input, r)
	slog.Info("Results", "part1", part1, "part2", part2)
}
