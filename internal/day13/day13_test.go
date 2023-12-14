package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func loadTestFile(id string) (input []string) {
	file, err := os.Open(filepath.Join("testdata", id+".txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func TestReadLavaMaps(t *testing.T) {
	tests := []struct {
		id    string
		input []string
		want  []lavaMap
	}{
		{
			id: "day13_test_input",
			want: []lavaMap{
				{"#.##..##.", "..#.##.#.", "##......#", "##......#", "..#.##.#.", "..##..##.", "#.#.##.#."},
				{"#...##..#", "#....#..#", "..##..###", "#####.##.", "#####.##.", "..##..###", "#....#..#"},
			},
		},
	}

	for _, tt := range tests {
		tt.input = loadTestFile(tt.id)
		fmt.Println(tt.input)
		t.Run(tt.id, func(t *testing.T) {
			got := readLavaMaps(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %#v, got %#v", tt.want, got)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	tests := []lavaMap{
		{"#.##..#", "..##...", "##..###", "#....#.", ".#..#.#", ".#..#.#", "#....#.", "##..###", "..##..."},
		{"##.##.#", "...##..", "..####.", "..####.", "#..##..", "##....#", "..####.", "..####.", "###..##"},
	}

	testMaps := readLavaMaps(loadTestFile("day13_test_input"))

	for idx, want := range tests {

		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			got := testMaps[idx].Rotate()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Expected %#v, got %#v", want, got)
			}
		})
	}
}

func TestCompareLines(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{"...##.", "...##.", 0},
		{"#.#.#..", "#.#.#.#", 1},
		{"#####", ".....", 5},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			got := compareLines(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Expected %#v, got %#v", tt.want, got)
			}
		})
	}
}

func TestPart1Handler(t *testing.T) {
	total, _ := part1Handler(loadTestFile("day13_test_input"))
	want := 405
	if total != want {
		t.Errorf("Expected %v, got %v", want, total)
	}
}

func TestPart2Handler(t *testing.T) {
	input := loadTestFile("day13_test_input")
	_, r := part1Handler(input)
	total := part2Handler(input, r)
	want := 400
	if total != want {
		t.Errorf("Expected %v, got %v", want, total)
	}
}
