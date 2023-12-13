package day12

import (
	"fmt"
	"testing"
)

func TestCheckSpringFit(t *testing.T) {
	tests := []struct {
		line   string
		spring int
		want   bool
	}{
		{".#?", 3, false},
		{"###", 3, true},
		{"?##", 3, true},
		{"####..", 3, false},
		{"##", 3, false},
		{"??##.", 3, false},
		{"?????#?..?#?", 2, true},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test index %d", i), func(t *testing.T) {
			got := checkSpringFit(tt.line, tt.spring)
			if got != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestNextSpringStart(t *testing.T) {
	tests := []struct {
		line string
		want struct {
			idx int
			eol bool
		}
	}{
		{".#?", struct {
			idx int
			eol bool
		}{1, false}},
		{"###", struct {
			idx int
			eol bool
		}{0, false}},
		{"...", struct {
			idx int
			eol bool
		}{0, true}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test index %d", i), func(t *testing.T) {
			idx, eol := nextSpringStart(tt.line)
			if idx != tt.want.idx || eol != tt.want.eol {
				t.Errorf("Expected (%v, %v), got (%v, %v)", tt.want.idx, tt.want.eol, idx, eol)
			}
		})
	}

}

func TestCountCombinations(t *testing.T) {
	tests := []struct {
		line         string
		springGroups []int
		combos       int
	}{
		{"?????#?..?#?", []int{2, 2, 2}, 10},
		{"??.?#??#?#??##????", []int{2, 4, 6, 1}, 2},
		{"???#???#.?#?????.#", []int{5, 1, 1, 1, 2, 1}, 2},
		{"???#???#.#??#??##??", []int{1, 1, 1, 1, 9}, 2},
		{"????###??.????#.#", []int{5, 1, 2, 1}, 7},
		{".?##?##?.?????#?#?.", []int{6, 6}, 4},
		{"?????#???#??#????", []int{3, 8}, 8},
		{"???#????.??", []int{3, 1, 1}, 13},
		{"?#?#???????#?#?..", []int{5, 4}, 4},
		{"??????##????.#??", []int{1, 9, 1, 1}, 3},
		{"#??????..???#??????", []int{2, 4, 1, 3, 1}, 11},
		{".#.???.#?#?#??##??.", []int{1, 1, 5, 4}, 6},
	}

	for i, tt := range tests {
		cache := make(map[string]int)
		t.Run(fmt.Sprintf("test index %d", i), func(t *testing.T) {
			got := countCombinations(tt.line, tt.springGroups, cache)
			if got != tt.combos {
				t.Errorf("Expected %v, got %v", tt.combos, got)
			}
		})
	}
}
