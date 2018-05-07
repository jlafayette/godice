package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {name string; expected int; x int; y int}{
		{"xmax", 3, 3, 2},
		{"same", 12, 12, 12},
		{"xneg", 0, -2, 0},
		{"2neg", -4, -4, -5},
		{"zeros", 0, 0, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Max(test.x, test.y)
			if test.expected != actual {
				t.Errorf("got %d, want %d", test.expected, actual)
			}
		})
	}
}


func TestMin(t *testing.T) {
	tests := []struct {name string; expected int; x int; y int}{
		{"ymin", 2, 3, 2},
		{"same", 12, 12, 12},
		{"xneg", -2, -2, 0},
		{"yneg", -5, -4, -5},
		{"zeros", 0, 0, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Min(test.x, test.y)
			if test.expected != actual {
				t.Errorf("got %d, want %d", test.expected, actual)
			}
		})
	}
}
