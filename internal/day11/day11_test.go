package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func LoadTestFile(id string) (input []string) {
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

func TestGalaxyLoad(t *testing.T) {
	input := LoadTestFile("day11_test_input")
	want := galaxy{
		coords: [][2]int{
			{0, 3}, {1, 7}, {2, 0}, {4, 6}, {5, 1},
			{6, 9}, {8, 7}, {9, 0}, {9, 4}},
		emptyRow:    map[int]bool{3: true, 7: true},
		emptyCol:    map[int]bool{2: true, 5: true, 8: true},
		emptyFactor: 2,
	}
	g := galaxy{}
	g.Load(input, 2)
	if !reflect.DeepEqual(g, want) {
		t.Errorf("wanted %v, got %v", want, g)
	}
}

func TestGalaxySteps(t *testing.T) {
	input := LoadTestFile("day11_test_input")
	tests := []struct {
		a     int
		b     int
		steps int
	}{
		{4, 8, 9},
		{0, 6, 15},
		{2, 5, 17},
		{7, 8, 5},
	}

	g := galaxy{}
	g.Load(input, 2)

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := g.Steps(g.coords[tt.a], g.coords[tt.b])
			if got != tt.steps {
				t.Errorf("wanted %v got %v\n", tt.steps, got)
			}
		})
	}
}
