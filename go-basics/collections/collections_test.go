package collections

import (
	"slices"
	"testing"
)

func TestRotateLeft(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		n    int
		want []int
	}{
		{"n=2", []int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{"n=7 wraps", []int{1, 2, 3, 4, 5}, 7, []int{3, 4, 5, 1, 2}},
		{"n=0", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		{"n=len", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"negative", []int{1, 2, 3, 4, 5}, -1, []int{5, 1, 2, 3, 4}},
		{"empty", []int{}, 3, []int{}},
		{"nil", nil, 1, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RotateLeft(tt.in, tt.n)
			if !slices.Equal(got, tt.want) {
				t.Errorf("RotateLeft(%v, %d) = %v, want %v", tt.in, tt.n, got, tt.want)
			}
		})
	}
}

func TestRotateLeftDoesNotModifyInput(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	RotateLeft(in, 2)
	if !slices.Equal(in, []int{1, 2, 3, 4, 5}) {
		t.Errorf("input was modified: %v", in)
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		in   []string
		want []string
	}{
		{[]string{"a", "b", "a", "c", "b"}, []string{"a", "b", "c"}},
		{[]string{"x", "x", "x"}, []string{"x"}},
		{[]string{}, []string{}},
	}
	for _, tt := range tests {
		got := Unique(tt.in)
		if !slices.Equal(got, tt.want) {
			t.Errorf("Unique(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
