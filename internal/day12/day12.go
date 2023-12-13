package day12

import (
	"bufio"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/foax/aoc2023/internal/util"
)

// check if a spring fits at the start of line
func checkSpringFit(line string, springCount int) bool {
	if len(line) < springCount {
		return false
	}
	for i, s := range line {
		if i == springCount {
			break
		}
		if s == '.' {
			return false
		}
	}
	if len(line) == springCount {
		return true
	}
	if line[springCount] == '#' {
		return false
	}
	return true
}

// Returns the index of where the next spring can start from.
// int: index; bool: true if hit the end of line.
func nextSpringStart(line string) (int, bool) {
	for i, s := range line {
		if s == '.' {
			continue
		}
		return i, false
	}
	return 0, true
}

// Count the possible combinations of springGroups in line, using results from cache to speed things up.
func countCombinations(line string, springGroups []int, cache map[string]int) int {
	slog.Debug("countCombinations start", "line", line, "springGroups", springGroups)
	defer slog.Debug("countCombinations end", "line", line, "springGroups", springGroups)

	// check the cache first; if we have a hit, return immediately
	if combo, hit := cache[fmt.Sprintf("%v", springGroups)+line]; hit {
		slog.Debug("countCombinations cache hit", "key", fmt.Sprintf("%v", springGroups)+line, "combo", combo)
		return combo
	}

	combos := 0
	j := 0
	i, eol := nextSpringStart(line)
	slog.Debug("countCombinations", "i", i, "eol", eol)
	if eol {
		return 0
	}

	for i < len(line) {
		slog.Debug("countCombinations in for loop", "line", line, "springGroups", springGroups, "i", i, "combos", combos, "remainingLine", line[i:])
		checkLine := line[i:]
		if len(checkLine) < util.SliceSum(springGroups)+len(springGroups)-1 {
			slog.Debug("countCombinations spring length greater than remaining length", "remaininglength", len(checkLine))
			break
		}

		// If any spring markers before the index are #, no further combinations are valid
		if strings.ContainsRune(line[0:i], '#') {
			slog.Debug("countCombinations # markers before string to check", "line", line[0:i])
			break
		}

		// If spring does not fit at the current index, increment index to the next valid starting point and check again
		if !checkSpringFit(checkLine, springGroups[0]) {
			slog.Debug("countCombinations spring does not fit", "line", checkLine, "spring", springGroups[0])
			j, eol = nextSpringStart(line[i+1:])
			if eol {
				break
			}
			i += j + 1
			continue
		}

		slog.Debug("countCombinations spring fits", "line", checkLine, "spring", springGroups[0])
		// j holds the index of where we want to start recursively checking for more combos from
		j = i + springGroups[0] + 1
		if line[i] == '#' {
			i += springGroups[0] + 1
		} else {
			i += 1
		}

		if len(springGroups[1:]) > 0 {
			if j >= len(line) {
				slog.Debug("countCombinations j > len(line), break", "springGroupsRemaining", springGroups[1:])
				break
			}
			newCombos := countCombinations(line[j:], springGroups[1:], cache)
			slog.Debug("countCombinations", "line", line, "springGroups", springGroups, "lineChecked", line[j:], "newCombos", newCombos)
			cacheKey := fmt.Sprintf("%v", springGroups[1:]) + line[j:]
			slog.Debug("countCombinations adding to cache", "cacheKey", cacheKey, "cacheValue", newCombos)
			cache[cacheKey] = newCombos
			combos += newCombos
		} else {
			if j < len(line) && (strings.ContainsRune(line[j:], '#')) {
				// no springs left to check but still # markers in remaining line.
				slog.Debug("countCombinations no springs left but # markers remain", "line", line)
				continue
			}
			combos++
		}
	}
	slog.Debug("countCombinations returning", "line", line, "springGroups", springGroups, "combos", combos)
	return combos
}

func parseInputLine(line string) (string, []int) {
	var springs []int
	fields := strings.Split(line, " ")
	for _, s := range strings.FieldsFunc(fields[1], func(r rune) bool { return r == ',' }) {
		x, _ := strconv.Atoi(s)
		springs = append(springs, x)
	}
	return fields[0], springs
}

func part1Handler(input []string, cache map[string]int) int {
	total := 0
	for _, line := range input {
		springLine, springs := parseInputLine(line)
		combo := countCombinations(springLine, springs, cache)
		slog.Debug("part1Handler", "springLine", springLine, "springs", springs, "combos", combo)
		total += combo
	}
	return total
}

func part2Handler(input []string, cache map[string]int) int {
	total := 0
	for _, line := range input {
		springLine, springs := parseInputLine(line)
		x := len(springLine)
		y := len(springs)
		for i := 1; i < 5; i++ {
			springLine = springLine + "?" + springLine[0:x]
			springs = append(springs, springs[0:y]...)
		}
		combo := countCombinations(springLine, springs, cache)
		slog.Debug("part1Handler", "springLine", springLine, "springs", springs, "combos", combo)
		total += combo
	}
	return total
}

func Execute(scanner *bufio.Scanner) {
	input := util.ReadInput(scanner)
	// share the cache between parts 1 and 2
	cache := make(map[string]int)
	part1 := part1Handler(input, cache)
	part2 := part2Handler(input, cache)
	slog.Info("Results", "part1", part1, "part2", part2, "cache size", len(cache))
}
