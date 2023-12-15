package day14

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/foax/aoc2023/internal/util"
)

// A vector of consequtive boulders
type boulderVector struct {
	start    int // what position this vector starts from
	length   int // the total length of the vectors
	boulders int // how many boulders were in this vector
}

type boulderPlatform [][]rune

// Create a new boulderPlatform from a list of strings
func NewFromString(input []string) (p boulderPlatform) {
	p = make([][]rune, len(input))
	for i := range input {
		p[i] = make([]rune, len(input[0]))
		for j := range input[i] {
			p[i][j] = rune(input[i][j])

		}
	}
	return p
}

// Convert a boulderPlatform into a list of []boulderVectors, one per column
func BoulderVectors(p boulderPlatform) (v [][]boulderVector) {
	// Easier to work with rows than columns
	pt := util.TransposeMatrix[rune](p)
	for i := range pt {
		row := make([]boulderVector, 0)
		vector := boulderVector{}
		for j := range pt[i] {
			if pt[i][j] == '#' {
				if vector != (boulderVector{}) {
					row = append(row, vector)
					vector = boulderVector{}
				}
				continue
			}
			if vector == (boulderVector{}) {
				vector = boulderVector{j, 0, 0}
			}
			vector.length++
			if pt[i][j] == 'O' {
				vector.boulders++
			}
		}
		if vector != (boulderVector{}) {
			row = append(row, vector)
		}
		v = append(v, row)
	}
	return
}

// Create a boulderPlatform from a list of []boulderVectors, reversing what BoulderVectors did
func NewFromVectors(v [][]boulderVector, platformLen int) (p boulderPlatform) {
	p = make(boulderPlatform, platformLen)
	for i := 0; i < platformLen; i++ {
		p[i] = make([]rune, len(v))
	}
	for j, vectors := range v {
		for _, vector := range vectors {
			for i := vector.start; i < vector.length+vector.start; i++ {
				switch {
				case i < vector.boulders+vector.start:
					p[i][j] = 'O'
				default:
					p[i][j] = '.'
				}
			}
		}
	}
	for i := range p {
		for j := range p[i] {
			if p[i][j] == rune(0) {
				p[i][j] = '#'
			}
		}
	}
	return p
}

func (p boulderPlatform) String() (s string) {
	for _, x := range p {
		s += fmt.Sprintln(string(x))
	}
	return
}

// Returns a north-tilted version of boulderPlatform
func TiltPlatform(p boulderPlatform) (tilted boulderPlatform) {
	// load = BoulderLoad(bv, len(pt[0]))
	tilted = NewFromVectors(BoulderVectors(p), len(p))
	return
}

// Cycles p through north, west, south, east
func CyclePlatform(p boulderPlatform) (cycled boulderPlatform) {
	for count := 1; count <= 4; count++ {
		p = TiltPlatform(p)
		p = util.RotateMatrix[rune](p, 1)
	}
	cycled = p
	return
}

// Calcuates the load on the north support
func BoulderLoad(p boulderPlatform) (load int) {
	for i, row := range p {
		l := len(p) - i
		for _, b := range row {
			if b == 'O' {
				load += l
			}
		}
	}
	return load
}

func part1Handler(input []string) int {
	p := NewFromString(input)
	p = TiltPlatform(p)
	return BoulderLoad(p)
}

func part2Handler(input []string) (total int) {
	p := NewFromString(input)
	turns := 1000000000
	// cache how many turns it took to reach a particular platform
	cache := make(map[string]int)
	// cache which platform was reached on a particular turn
	reverseCache := make(map[int]string)
	for turn := 1; turn <= turns; turn++ {
		p = CyclePlatform(p)
		id := fmt.Sprintf("%v", p)
		_, ok := cache[id]
		if ok { // cache hit
			// find which turn the next platform was found on
			p = CyclePlatform(p)
			nextFound := cache[fmt.Sprintf("%v", p)]
			// calculate which existing turn had the answer's platform
			answerTurn := nextFound + ((turns-(nextFound-1))%(turn-nextFound+1) - 1)
			// calculate the load of the platform
			answerString := strings.TrimSuffix(reverseCache[answerTurn], "\n")
			answerBoulder := NewFromString(strings.Split(answerString, "\n"))
			answer := BoulderLoad(answerBoulder)
			return answer
		}
		cache[id] = turn
		reverseCache[turn] = id
	}
	return 0
}

func Execute(scanner *bufio.Scanner) {
	input := util.ReadInput(scanner)
	part1 := part1Handler(input)
	part2 := part2Handler(input)
	fmt.Println(part1, part2)
}
