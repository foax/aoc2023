package day14

import (
	"bufio"
	"fmt"

	"github.com/foax/aoc2023/internal/util"
)

type boulderVector struct {
	start    int
	length   int
	boulders int
}

type boulderPlatform [][]boulderVector

func (p *boulderPlatform) Load(input []string) {
	// *p = make(boulderPlatform, len(input[0]))
	for j := range input[0] {
		var vector boulderVector
		row := make([]boulderVector, 0)
		for i := range input {
			if input[i][j] == '#' {
				if vector != (boulderVector{}) {
					row = append(row, vector)
					vector = boulderVector{}
				}
				continue
			}
			if vector == (boulderVector{}) {
				vector = boulderVector{i, 0, 0}
			}
			vector.length++
			if input[i][j] == 'O' {
				vector.boulders++
			}
		}
		if vector != (boulderVector{}) {
			row = append(row, vector)
		}
		*p = append(*p, row)
	}
}

func part1Handler(p boulderPlatform, platformLen int) (total int) {
	for _, row := range p {
		for _, vector := range row {
			total += util.SumOfDescendingSequence(platformLen-vector.start, vector.boulders)
		}
	}
	return
}

func Execute(scanner *bufio.Scanner) {
	var p boulderPlatform
	input := util.ReadInput(scanner)
	p.Load(input)
	part1 := part1Handler(p, len(input))
	fmt.Println(part1)

}
