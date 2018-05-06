package main

import (
	"fmt"
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
		t.Errorf("Test failed", expected20, actual20)
	}
	expected6 := []int{1, 2, 3, 4, 5, 6}
	actual6 := DS(6)
	if !(testEq(actual6, expected6)) {
		t.Errorf("Test failed", expected6, actual6)
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
			return fmt.Sprintf("Unequal Values", compare[i], v)
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
		} else {
			t.Logf("Passed test for %d", n)
		}
	}
}
