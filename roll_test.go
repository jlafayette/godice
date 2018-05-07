package main

import (
	"fmt"
	"sort"
	"testing"
)

func testEq(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestDS(t *testing.T) {
	expected20 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	actual20 := DS(20)
	if !(testEq(actual20, expected20)) {
		t.Errorf("Expected %d got %d", expected20, actual20)
	}
	expected6 := []int{1, 2, 3, 4, 5, 6}
	actual6 := DS(6)
	if !(testEq(actual6, expected6)) {
		t.Errorf("Expected %d got %d", expected6, actual6)
	}
}

// Use results of DS to test DC since they should match.
func testDC(sides int) string {
	compare := DS(sides)
	i := 0
	for v := range DC(sides) {
		if i >= len(compare) {
			return fmt.Sprintf("Too many results, expected %d, got %d+", len(compare), i+1)
			break
		}
		if v != compare[i] {
			return fmt.Sprintf("Unequal Values %d %d", compare[i], v)
		}
		i++
	}
	if i < len(compare) {
		return fmt.Sprintf("Not enough results, expected %d, got %d", len(compare), i)
	}
	return ""
}

func TestDC(t *testing.T) {
	for _, n := range [...]int{4, 6, 8, 10, 12, 20} {
		msg := testDC(n)
		if msg != "" {
			t.Errorf(msg)
		}
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
