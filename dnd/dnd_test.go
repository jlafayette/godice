package dnd

import (
	"fmt"
	"math/big"
	"reflect"
	"sort"
	"testing"
)

type rollTest struct {
	roll          DnDRoll
	expDMap       map[int]int
	expAverage    float64
	expAverageRat *big.Rat
}

func (r *rollTest) Run(t *testing.T) {
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

func TestD20(t *testing.T) {
	rt := rollTest{
		D20(),
		map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1, 8: 1, 9: 1, 10: 1, 11: 1, 12: 1, 13: 1, 14: 1, 15: 1, 16: 1, 17: 1, 18: 1, 19: 1, 20: 1},
		10.5,
		big.NewRat(21, 2),
	}
	rt.Run(t)
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

func TestDisadvantage(t *testing.T) {
	rt := rollTest{
		Disadvantage(),
		map[int]int{1: 39, 2: 37, 3: 35, 4: 33, 5: 31, 6: 29, 7: 27, 8: 25, 9: 23, 10: 21, 11: 19, 12: 17, 13: 15, 14: 13, 15: 11, 16: 9, 17: 7, 18: 5, 19: 3, 20: 1},
		7.175,
		big.NewRat(287, 40),
	}
	rt.Run(t)
}

func TestStat(t *testing.T) {
	rt := rollTest{
		Stat(),
		map[int]int{3: 1, 4: 4, 5: 10, 6: 21, 7: 38, 8: 62, 9: 91, 10: 122, 11: 148, 12: 167, 13: 172, 14: 160, 15: 131, 16: 94, 17: 54, 18: 21},
		12.244598765432098,
		big.NewRat(15869, 1296),
	}
	rt.Run(t)
}

func TestCoryStat(t *testing.T) {
	rt := rollTest{
		CoryStat(),
		map[int]int{8: 1, 9: 2, 10: 3, 11: 4, 12: 5, 13: 6, 14: 5, 15: 4, 16: 3, 17: 2, 18: 1},
		13.0,
		big.NewRat(13, 1),
	}
	rt.Run(t)
}

func TestBattleAxe(t *testing.T) {
	rt := rollTest{
		BattleAxe(),
		map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1, 8: 1, 9: 1, 10: 1, 11: 1, 12: 1},
		6.5,
		big.NewRat(13, 2),
	}
	rt.Run(t)
}

func TestMaul(t *testing.T) {
	rt := rollTest{
		Maul(),
		map[int]int{2: 1, 3: 2, 4: 3, 5: 4, 6: 5, 7: 6, 8: 5, 9: 4, 10: 3, 11: 2, 12: 1},
		7,
		big.NewRat(7, 1),
	}
	rt.Run(t)
}

//func Test1(t *testing.T) {
//	generateDMapData(CoryStat())
//}

// Print sorted DMap of a roll
func generateDMapData(a DnDRoll) {
	m := a.DMap()
	keys := make([]int, len(m))
	ki := 0
	for k := range m {
		keys[ki] = k
		ki++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	fmt.Print("\nmap[int]int{")
	for i, v := range keys {
		if i == len(keys)-1 {
			fmt.Printf("%d:%d", v, m[v])
		} else {
			fmt.Printf("%d:%d, ", v, m[v])
		}
	}
	fmt.Print("}\n")
}

// Useful for validating map against percentage output
func printPercentages(a DnDRoll) {
	m := a.DMap()
	fmt.Print("\n #   %")
	for i := 1; i < 21; i++ {
		r := big.NewRat(int64(m[i]*100), int64(400))
		fmt.Printf("\n%2d %s", i, r.FloatString(2))
	}
}
