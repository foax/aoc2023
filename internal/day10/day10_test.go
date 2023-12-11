package day10

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func readInputFile(id string) (grid [][]rune) {
	file, err := os.Open(filepath.Join("testdata", id+".txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func TestFindStart(t *testing.T) {
	cases := []struct {
		id   string
		grid [][]rune
		want [2]int
	}{
		{id: "test_input_1", want: [2]int{1, 1}},
		{id: "test_input_2", want: [2]int{2, 0}},
		{id: "test_input_3", want: [2]int{1, 1}},
		{id: "test_input_4", want: [2]int{4, 12}},
		{id: "test_input_5", want: [2]int{0, 4}},
	}

	for idx := range cases {
		cases[idx].grid = readInputFile(cases[idx].id)
	}

	for _, c := range cases {
		t.Run(c.id, func(t *testing.T) {
			got := findStart(c.grid)
			if c.want != got {
				t.Errorf("Expected %v, got %v", c.want, got)
			}
		})
	}
}

func TestLoopLength(t *testing.T) {
	cases := []struct {
		id   string
		grid [][]rune
		want int
	}{
		{id: "test_input_1", want: 8},
		{id: "test_input_2", want: 16},
		{id: "test_input_3", want: 46},
		{id: "test_input_4", want: 140},
		{id: "test_input_5", want: 160},
	}

	for idx := range cases {
		cases[idx].grid = readInputFile(cases[idx].id)
	}

	for _, c := range cases {
		t.Run(c.id, func(t *testing.T) {
			got, _ := loopLength(c.grid, findStart(c.grid))
			if c.want != got {
				t.Errorf("Expected %v, got %v", c.want, got)
			}
		})
	}
}
