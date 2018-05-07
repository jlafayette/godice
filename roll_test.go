package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestDS(t *testing.T) {
	tests := []struct {
		name string
		in int
		out [] int
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
		in int
		out [] int
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

// Test if values match (unordered)
// WARNING: Has side effect of sorting both of the input slices
func testUnorderedEq(a []int, b []int) string {
	if len(a) != len(b) {
		return fmt.Sprintf("Unequal lengths %d != %d", len(a), len(b))
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	//fmt.Println("a", a)
	//fmt.Println("b", b)
	for idx := range a {
		if a[idx] != b[idx] {
			return fmt.Sprintf("Sorted values at index %d do not match %d != %d", idx, a[idx], b[idx])
		}
	}
	return ""
}

// Convert chanel to slice
func testUnorderedEqWrapper(s []int, c <-chan int) string {
	converted := make([]int, 0, len(s))
	for v := range c {
		converted = append(converted, v)
	}
	return testUnorderedEq(s, converted)
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

func TestDR2(t *testing.T) {
	expected := []int{2, 3, 4, 5, 6, 7, 3, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 9, 5, 6, 7, 8, 9, 10, 6, 7, 8, 9, 10, 11, 7, 8, 9, 10, 11, 12}
	msg := testUnorderedEqWrapper(expected, DR2(defaultSum, 6, 6))
	if msg != "" {
		t.Errorf(msg)
	}
	ex4x4 := []int{10, 5, 6, 7, 8, 6, 7, 8, 9, 7, 5, 4, 6, 7, 5, 6, 7, 8, 6, 7, 8, 9, 7, 8, 9, 8, 9, 10, 8, 9, 10, 11, 6, 7, 8, 9, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 5, 6, 7, 8, 6, 7, 8, 9, 7, 8, 9, 10, 8, 9, 10, 11, 6, 7, 8, 9, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 11, 12, 13, 14, 6, 7, 8, 9, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 11, 12, 13, 14, 9, 10, 11, 12, 10, 11, 12, 13, 11, 12, 13, 14, 12, 13, 14, 15, 7, 8, 9, 10, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 8, 9, 10, 11, 9, 10, 11, 12, 10, 11, 12, 13, 11, 12, 13, 14, 9, 10, 11, 12, 10, 11, 12, 13, 11, 12, 13, 14, 12, 13, 14, 15, 10, 11, 12, 13, 11, 12, 13, 14, 12, 13, 14, 15, 13, 14, 15, 16}
	msg = testUnorderedEqWrapper(ex4x4, DR2(defaultSum, 4, 4, 4, 4))
	if msg != "" {
		t.Errorf(msg)
	}
}
