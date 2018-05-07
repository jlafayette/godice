package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	test := func(expected, x, y int) {
		actual := Max(x, y)
		if expected != actual {
			t.Errorf("Test failed. Expected %d got %d", expected, actual)
		}
	}
	test(3, 3, 2)
	test(12, 12, 12)
	test(0, -2, 0)
	test(-4, -4, -5)
	test(0, 0, 0)
}

func TestMin(t *testing.T) {
	test := func(expected, x, y int) {
		actual := Min(x, y)
		if expected != actual {
			t.Errorf("Test failed. Expected %d got %d", expected, actual)
		}
	}
	test(2, 3, 2)
	test(12, 12, 12)
	test(-2, -2, 0)
	test(-5, -4, -5)
	test(0, 0, 0)
}
