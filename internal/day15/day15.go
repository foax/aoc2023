package day15

import (
	"bufio"
	"fmt"
	"log/slog"
	"strings"

	"github.com/foax/aoc2023/internal/util"
)

type Box struct {
	firstSlot *LensSlot
	lenses    map[string]*LensSlot
}

type Lens struct {
	label       string
	focalLength int
}

// linked list ftw
type LensSlot struct {
	lens Lens
	next *LensSlot
}

type Step struct {
	label string
	box   int
	op    rune
	focal int
}

func Hash(s string) (value int) {
	for _, c := range s {
		value += int(c)
		value *= 17
		value = value % 256
	}
	return
}

func parseStep(input string) (s Step) {
	for _, c := range input {
		switch {
		case c == '=' || c == '-':
			s.op = c
		case s.op == rune('='):
			s.focal = s.focal*10 + int(c) - int('0')
		default:
			s.label += string(c)
		}
	}
	s.box = Hash(s.label)
	return
}

func printBox(b Box, i int) {
	fmt.Printf("Box %v: ", i)
	for j, lensSlot := 1, b.firstSlot; lensSlot != nil; j++ {
		fmt.Printf("[%v %d] ", lensSlot.lens.label, lensSlot.lens.focalLength)
		lensSlot = lensSlot.next
	}
	fmt.Printf("\n")
}

func part2Handler(input string) (total int) {
	var boxes [256]Box
	for i := range boxes {
		boxes[i].lenses = make(map[string]*LensSlot)
	}

	for _, stepStr := range strings.Split(input, ",") {
		step := parseStep(stepStr)
		switch step.op {
		case '-':
			if slot, ok := boxes[step.box].lenses[step.label]; ok {
				if boxes[step.box].firstSlot == slot {
					boxes[step.box].firstSlot = slot.next
				} else {
					prevSlot := boxes[step.box].firstSlot
					for prevSlot.next != slot {
						prevSlot = prevSlot.next
					}
					prevSlot.next = prevSlot.next.next
				}
				delete(boxes[step.box].lenses, step.label)
			}
		case '=':
			if slot, ok := boxes[step.box].lenses[step.label]; ok {
				slot.lens.focalLength = step.focal
			} else {
				slot := LensSlot{lens: Lens{step.label, step.focal}}
				if boxes[step.box].firstSlot == nil {
					boxes[step.box].firstSlot = &slot
				} else {
					lastSlot := boxes[step.box].firstSlot
					for lastSlot.next != nil {
						lastSlot = lastSlot.next
					}
					lastSlot.next = &slot
				}
				boxes[step.box].lenses[step.label] = &slot
			}
		}
	}

	for i := range boxes {
		for j, lensSlot := 1, boxes[i].firstSlot; lensSlot != nil; j++ {
			total += (i + 1) * j * lensSlot.lens.focalLength
			lensSlot = lensSlot.next
		}
	}
	return
}

func part1Handler(input string) (total int) {
	for _, s := range strings.Split(input, ",") {
		total += Hash(s)
	}
	return
}

func Execute(scanner *bufio.Scanner) {
	input := util.ReadInput(scanner)
	part1 := part1Handler(input[0])
	part2 := part2Handler(input[0])
	slog.Info("Results", "part1", part1, "part2", part2)

}
