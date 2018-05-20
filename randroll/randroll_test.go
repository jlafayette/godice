package randroll

import (
	"github.com/jlafayette/godice/roll"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkR(b *testing.B) {
	benchmarks := []struct {
		name string
		dice []int
	}{
		{"d20", []int{20}},
		{"2d20", []int{20, 20}},
		{"3d6", []int{6, 6, 6}},
		{"3d2", []int{2, 2, 2}},
		{"10d10", []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}},
	}

	time.Sleep(time.Nanosecond)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	b.ResetTimer()

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, d := range bm.dice {
					R(d, *r)
				}
			}
		})
	}
}

func BenchmarkRD(b *testing.B) {
	benchmarks := []struct {
		name string
		dice []int
	}{
		{"d20", []int{20}},
		{"2d20", []int{20, 20}},
		{"3d6", []int{6, 6, 6}},
		{"3d2", []int{2, 2, 2}},
		{"10d10", []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}},
	}

	time.Sleep(time.Nanosecond)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	b.ResetTimer()

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RD(bm.dice, *r)
			}
		})
	}
}

func TestRandAverage(t *testing.T) {
	tests := []struct {
		name     string
		dice     []int
		sumfn    sumFn
		expected *big.Rat
	}{
		{"d20", []int{20}, roll.DefaultSum, big.NewRat(21, 2)},
		{"Advantage", []int{20, 20}, roll.DropLowest, big.NewRat(553, 40)},
		{"Disadvantage", []int{20, 20}, roll.DropHighest, big.NewRat(287, 40)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := RandAverage(test.dice, test.sumfn, 1000000, 20)
			t.Logf("Got %s, expected %s", r.FloatString(3), test.expected.FloatString(3))
		})
	}
}

func TestRMap(t *testing.T) {
	count := 1000000
	sides := 2

	time.Sleep(time.Nanosecond)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	m := make(map[int]int)
	for i := 0; i < count; i++ {
		result := R(sides, *r)
		m[result] = m[result] + 1
	}
	t.Log(m)
}
