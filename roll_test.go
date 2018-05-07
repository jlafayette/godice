package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestDS(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  []int
	}{
		{"d4", 4, []int{1, 2, 3, 4}},
		{"d6", 6, []int{1, 2, 3, 4, 5, 6}},
		{"d8", 8, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"d10", 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"d12", 12, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
		{"d20", 20, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := DS(test.in)

			if len(out) != len(test.out) {
				t.Errorf("Wrong length, got %d, want %d", len(out), len(test.out))
			}
			for i := range out {
				if out[i] != test.out[i] {
					t.Errorf("Wrong value at index %d, got %d, want %d", i, out[i], test.out[i])
				}
			}
		})
	}
}

func TestDC(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  []int
	}{
		{"d4", 4, []int{1, 2, 3, 4}},
		{"d6", 6, []int{1, 2, 3, 4, 5, 6}},
		{"d8", 8, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"d10", 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"d12", 12, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
		{"d20", 20, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := DC(test.in)
			i := 0
			for v := range out {
				if i >= len(test.out) {
					t.Errorf("Too many results, got %d+, want %d", i+1, len(test.out))
					break
				}
				if v != test.out[i] {
					t.Errorf("Unequal Values, got %d, want %d", v, test.out[i])
				}
				i++
			}
			if i < len(test.out) {
				t.Errorf("Not enough results, got %d, want %d", i, len(test.out))
			}
		})
	}
}

func TestDR2(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{"1d12", []int{12}, []int{7, 11, 12, 8, 9, 1, 2, 3, 4, 5, 6, 10}},
		{"2d6", []int{6, 6}, []int{2, 3, 4, 5, 6, 7, 3, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 9, 5, 6, 7, 8, 9, 10, 6, 7, 8, 9, 10, 11, 7, 8, 9, 10, 11, 12}},
		{"4d4", []int{4, 4, 4, 4}, []int{10, 5, 6, 7, 8, 6, 7, 8, 9, 7, 5, 4, 6, 7, 5, 6, 7, 8, 6, 7, 8, 9, 7, 8, 9, 8, 9, 10, 8, 9, 10, 11, 6, 7, 8, 9, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 5, 6, 7, 8, 6, 7, 8, 9, 7, 8, 9, 10, 8, 9, 10, 11, 6, 7, 8, 9, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 11, 12, 13, 14, 6, 7, 8, 9, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 11, 12, 13, 14, 9, 10, 11, 12, 10, 11, 12, 13, 11, 12, 13, 14, 12, 13, 14, 15, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 11, 12, 13, 14, 9, 10, 11, 12, 10, 11, 12, 13, 11, 12, 13, 14, 12, 13, 14, 15, 10, 11, 12, 13, 11, 12, 13, 14, 12, 13, 14, 15, 13, 14, 15, 16}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := DR2(defaultSum, test.in...)
			out := make([]int, 0, len(test.out))
			for v := range c {
				out = append(out, v)
			}
			if len(out) != len(test.out) {
				t.Errorf("Got length of %d, expected %d", len(out), len(test.out))
				return
			}
			sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
			sort.Slice(test.out, func(i, j int) bool { return test.out[i] < test.out[j] })
			for idx := range out {
				if out[idx] != test.out[idx] {
					t.Errorf("Wrong value at index %d, got %d, expected %d", idx, out[idx], test.out[idx])
				}
			}
		})
	}
}

func TestDMap(t *testing.T) {
	tests := []struct {
		name  string
		sumfn sumFn
		dice  []int
		out   map[int]int
	}{
		{"1d4", defaultSum, []int{4}, map[int]int{1: 1, 2: 1, 3: 1, 4: 1}},
		{"2d6", defaultSum, []int{6, 6}, map[int]int{2: 1, 3: 2, 4: 3, 5: 4, 6: 5, 7: 6, 8: 5, 9: 4, 10: 3, 11: 2, 12: 1}},
		{"3d6", defaultSum, []int{6, 6, 6}, map[int]int{3: 1, 9: 25, 4: 3, 5: 6, 7: 15, 10: 27, 12: 25, 16: 6, 18: 1, 8: 21, 11: 27, 13: 21, 17: 3, 6: 10, 14: 15, 15: 10}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := DMap(test.sumfn, test.dice...)
			eq := reflect.DeepEqual(out, test.out)
			if !eq {
				t.Errorf("Got %d, expected %d", out, test.out)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name    string
		sumfn   sumFn
		dice    []int
		average float64
	}{
		{"d4", defaultSum, []int{4}, 2.5},
		{"d6", defaultSum, []int{6}, 3.5},
		{"d8", defaultSum, []int{8}, 4.5},
		{"d10", defaultSum, []int{10}, 5.5},
		{"d12", defaultSum, []int{12}, 6.5},
		{"d20", defaultSum, []int{20}, 10.5},
		{"2d6", defaultSum, []int{6, 6}, 7},
		{"Advantage", dropLowest, []int{20, 20}, 13.825},
		{"Disadvantage", dropHighest, []int{20, 20}, 7.175},
		{"3d6", defaultSum, []int{6, 6, 6}, 10.5},
		{"4d6 drop lowest", dropLowest, []int{6, 6, 6, 6}, 12.244598765432098},
		{"1d12", defaultSum, []int{12}, 6.5},
		{"1d12 reroll 1&2", rerollBelow2, []int{12, 12}, 7.333333333333333},
		{"1d4 reroll 1&2", rerollBelow2, []int{4}, 2.5},  // doesn't work on 1 dice
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			average := Average(test.sumfn, test.dice...)
			if average != test.average {
				t.Errorf("Got %f, expected %f", average, test.average)
			}
		})
	}
}

// Useful for generating output to test against DR2 (obviously it should be validated before using to test)
func generateTestData() {
	fmt.Println()
	i := 0
	for v := range DR2(defaultSum, 4, 4, 4, 4) {
		fmt.Print(v, ", ")
		i++
	}
	fmt.Println("\ni:", i)
}

func generateDMapData() {
	fmt.Println()
	m := DMap(defaultSum, 6, 6, 6)
	for k, v := range m {
		fmt.Printf("%d: %d, ", k, v)
	}
}
