package day09

import (
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	logger := slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{}))
	slog.SetDefault(logger)
	m.Run()
}

func TestLoadDataset(t *testing.T) {
	input := []string{"0 3 6 9 12 15", "1 3 6 10 15 21", "10 13 16 21 30 45"}
	want := [][]int{{0, 3, 6, 9, 12, 15}, {1, 3, 6, 10, 15, 21}, {10, 13, 16, 21, 30, 45}}
	got := loadDataset(input)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("loadDataset failed; wanted %v got %v\n", want, got)
	}
}

func TestNextSequence(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{0, 3, 6, 9, 12, 15}, []int{3, 3, 3, 3, 3}},
		{[]int{1, 3, 6, 10, 15, 21}, []int{2, 3, 4, 5, 6}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := nextSequence(tt.input)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("nextSequence failed: wanted %v got %v\n", tt.want, got)
			}

		})
	}
}
