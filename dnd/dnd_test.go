package dnd

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

type rollTest struct {
	roll DnDRoll
	expDMap map[int]int
	expAverage float64
	expAverageRat *big.Rat
}

func (r *rollTest) Run (t *testing.T) {
	t.Run("DMap", func(t *testing.T) {
		actual := r.roll.DMap()
		eq := reflect.DeepEqual(actual, r.expDMap)
		if !eq {
			t.Errorf("Got %d, expected %d", actual, r.expDMap)
		}
	})
	t.Run("Average", func(t *testing.T) {
		actual := r.roll.Average()
		if r.expAverage != actual {
			t.Errorf("Got %f, expected %f", actual, r.expAverage)
		}
	})
	t.Run("AverageRat", func(t *testing.T) {
		actual := r.roll.AverageRat()
		if r.expAverageRat.Cmp(actual) != 0 {
			t.Errorf("Got %s, expected %s", actual.String(), r.expAverageRat.String())
		}
	})
	t.Run("Rand", func(t *testing.T) {
		t.Log(r.roll.Rand(10))
	})
}

func TestAdvantage(t *testing.T) {
	rt := rollTest{
		Advantage(),
		map[int]int{1: 1, 2: 3, 3: 5, 4: 7, 5: 9, 6: 11, 7: 13, 8: 15, 9: 17, 10: 19, 11: 21, 12: 23, 13: 25, 14: 27, 15: 29, 16: 31, 17: 33, 18: 35, 19: 37, 20: 39},
		13.825,
		big.NewRat(553, 40),
	}
	rt.Run(t)
}

// Useful for generating test data for maps
func generateDMapData(a DnDRoll) {
	m := a.DMap()
	fmt.Print("\nmap[int]int{")
	for i := 1; i < a.dice[0]+1; i++ {
		fmt.Printf("%d:%d, ", i, m[i])
	}
	fmt.Print("}\n")
}

// Useful for validating map against percentage output
func printPercentages() {
	a := Advantage()
	m := a.DMap()
	fmt.Print("\n #   %")
	for i := 1; i < 21; i++ {
		r := big.NewRat(int64(m[i]*100), int64(400))
		fmt.Printf("\n%2d %s", i, r.FloatString(2))
	}
}
