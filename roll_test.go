package main

import (
	"fmt"
	"math/big"
	"reflect"
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

func TestExplode(t *testing.T) {
	tests := []struct {
		name   string
		reroll bool
		in     []int
		out    [][]int
	}{
		{"d2", false, []int{2}, [][]int{{1}, {2}}},
		{"d2reroll", true, []int{2}, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := Explode(test.reroll, test.in...)
			var out [][]int
			for v := range c {
				out = append(out, v)
			}
			if len(out) != len(test.out) {
				t.Errorf("Got length of %d, expected %d", len(out), len(test.out))
				return
			}
			eq := reflect.DeepEqual(out, test.out)
			if !eq {
				t.Errorf("Got %d, expected %d", out, test.out)
			}
		})
	}
}

func TestDMap(t *testing.T) {
	tests := []struct {
		name   string
		reroll bool
		sumfn  sumFn
		dice   []int
		out    map[int]int
	}{
		{"1d4", false, defaultSum, []int{4}, map[int]int{1: 1, 2: 1, 3: 1, 4: 1}},
		{"2d6", false, defaultSum, []int{6, 6}, map[int]int{2: 1, 3: 2, 4: 3, 5: 4, 6: 5, 7: 6, 8: 5, 9: 4, 10: 3, 11: 2, 12: 1}},
		{"3d6", false, defaultSum, []int{6, 6, 6}, map[int]int{3: 1, 9: 25, 4: 3, 5: 6, 7: 15, 10: 27, 12: 25, 16: 6, 18: 1, 8: 21, 11: 27, 13: 21, 17: 3, 6: 10, 14: 15, 15: 10}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := DMap(test.reroll, test.sumfn, test.dice...)
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
		reroll  bool
		sumfn   sumFn
		dice    []int
		average float64
	}{
		{"d4", false, defaultSum, []int{4}, 2.5},
		{"d6", false, defaultSum, []int{6}, 3.5},
		{"d8", false, defaultSum, []int{8}, 4.5},
		{"d10", false, defaultSum, []int{10}, 5.5},
		{"d12", false, defaultSum, []int{12}, 6.5},
		{"d20", false, defaultSum, []int{20}, 10.5},
		{"2d6", false, defaultSum, []int{6, 6}, 7},
		{"Advantage", true, dropLowest, []int{20}, 13.825},
		{"Disadvantage", true, dropHighest, []int{20}, 7.175},
		{"3d6", false, defaultSum, []int{6, 6, 6}, 10.5},
		{"4d6", false, defaultSum, []int{6, 6, 6, 6}, 14},
		{"4d2", false, defaultSum, []int{2, 2, 2, 2}, 6},
		{"4d6 drop lowest", false, dropLowest, []int{6, 6, 6, 6}, 12.244598765432098},
		{"1d12", false, defaultSum, []int{12}, 6.5},
		{"1d12 reroll 1&2", true, rerollOneAndTwo, []int{12}, 7.333333333333333},
		{"1d4 reroll 1&2", true, rerollOneAndTwo, []int{4}, 3.0},
	}
	//12.244598765428275 (from AnyDice)
	//12.244598765432098

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			average := Average(test.reroll, test.sumfn, test.dice...)
			if average != test.average {
				t.Errorf("Got %f, expected %f", average, test.average)
			}
		})
	}
}

func TestAverageF(t *testing.T) {
	tests := []struct {
		name    string
		reroll  bool
		sumfn   sumFn
		dice    []int
		average *big.Rat
	}{
		{"d4", false, defaultSum, []int{4}, big.NewRat(5, 2)},
		{"d6", false, defaultSum, []int{6}, big.NewRat(7, 2)},
		{"d8", false, defaultSum, []int{8}, big.NewRat(9, 2)},
		{"d10", false, defaultSum, []int{10}, big.NewRat(11, 2)},
		{"d12", false, defaultSum, []int{12}, big.NewRat(13, 2)},
		{"d20", false, defaultSum, []int{20}, big.NewRat(21, 2)},
		{"2d6", false, defaultSum, []int{6, 6}, big.NewRat(7, 1)},
		{"Advantage", true, dropLowest, []int{20}, big.NewRat(553, 40)},
		{"Disadvantage", true, dropHighest, []int{20}, big.NewRat(287, 40)},
		{"3d6", false, defaultSum, []int{6, 6, 6}, big.NewRat(21, 2)},
		{"4d6 drop lowest", false, dropLowest, []int{6, 6, 6, 6}, big.NewRat(15869, 1296)},
		{"1d12", false, defaultSum, []int{12}, big.NewRat(13, 2)},
		{"1d12 reroll 1&2", true, rerollOneAndTwo, []int{12}, big.NewRat(22, 3)},
		{"1d4 reroll 1&2", true, rerollOneAndTwo, []int{4}, big.NewRat(3, 1)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			average := AverageRat(test.reroll, test.sumfn, test.dice...)
			if average.Cmp(test.average) != 0 {
				t.Errorf("Got %s, expected %s", average.String(), test.average.String())
				t.Logf("Result as fraction: %s", average.FloatString(16))
			}
		})
	}
}

func generateDMapData() {
	fmt.Println()
	m := DMap(false, defaultSum, 6, 6, 6)
	for k, v := range m {
		fmt.Printf("%d: %d, ", k, v)
	}
}

func closeEnough(ep float64, a, b float64) bool {
	if (a-b) < ep && (b-a) < ep {
		return true
	}
	return false
}

func findRat() {
	//7.175
	// 12.244598765432098
	// 7.333333333333333
	whole := int64(12)
	tgt := 0.244598765432098
	base := int64(6 * 6 * 6 * 6)
	f := big.NewRat(1, base)
	for i := int64(2); i < base; i++ {
		f.SetFrac64(i, base)
		n, _ := f.Float64()
		//fmt.Print("R:", f, " f:",  n, " ")

		if closeEnough(0.00000000000001, n, tgt) {
			fmt.Println()
			fmt.Println(f, "==", tgt)

			num := whole*base + i
			f.SetFrac64(num, base)
			n2, _ := f.Float64()
			fmt.Println("Rat:", f, " Float:", n2)

			break
		}
		if i%10 == 0 {
			//fmt.Println()
		}
	}
}
