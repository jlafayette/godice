package dnd

import (
	"github.com/jlafayette/godice/randroll"
	"github.com/jlafayette/godice/roll"
	"math/big"
	"math/rand"
	"time"
)

type DnDRoll struct {
	reroll   bool
	sumfn    roll.SumFn
	dice     []int
	modifier int
}

func (d *DnDRoll) sumMod(r []int) int {
	return d.sumfn(r) + d.modifier
}
func (d *DnDRoll) DMap() map[int]int {
	return roll.DMap(d.reroll, d.sumMod, d.dice...)
}
func (d *DnDRoll) Average() float64 {
	return roll.Average(d.reroll, d.sumMod, d.dice...)
}
func (d *DnDRoll) AverageRat() *big.Rat {
	return roll.AverageRat(d.reroll, d.sumMod, d.dice...)
}
func (d *DnDRoll) Rand(n int) []int {
	time.Sleep(time.Nanosecond)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	sums := make([]int, n)
	for si := range sums {
		rolls := make([]int, len(d.dice))
		for di, dice := range d.dice {
			rolls[di] = randroll.R(dice, *r)
		}
		sums[si] = d.sumMod(rolls)
	}
	return sums
}

func D20() DnDRoll {
	return DnDRoll{false, roll.DefaultSum, []int{20}, 0}
}

func Advantage() DnDRoll {
	return DnDRoll{true, roll.DropLowest, []int{20}, 0}
}

func Disadvantage() DnDRoll {
	return DnDRoll{true, roll.DropHighest, []int{20}, 0}
}

func Stat() DnDRoll {
	return DnDRoll{false, roll.DropLowest, []int{6, 6, 6, 6}, 0}
}

func CoryStat() DnDRoll {
	return DnDRoll{false, roll.DefaultSum, []int{6, 6}, 6}
}

func BattleAxe() DnDRoll {
	return DnDRoll{false, roll.DefaultSum, []int{12}, 0}
}

func Maul() DnDRoll {
	return DnDRoll{false, roll.DefaultSum, []int{6, 6}, 0}
}
