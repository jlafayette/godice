// sum			   (rolls)
// reroll take 2nd (range|condition, rolls)
// keep x highest  (int, rolls)
// keep x lowest   (int, rolls)

/* USES
Advantage    keep 1 highest
	Highest(1, 20,20)
Disadvantage keep 1 lowest
	Lowest(1, 20,20)
Character Creation keep 3 highest of 4
	Highest(3, 6,6,6,6)
Damage
	Roll(12)
	Roll(6,6,6,6)
Damage + reroll 1&2, keep 2nd
	6,6
Roll + constant
Roll - constant

Dice Interface
	Sum ([]int) int
	Explode (...int)

For each roll:
Frequency
	FMap(dice{}) map[int]int
Frequency vs target
	FTgtMap(dice{})
Random series of rolls
	RR(dice{}) []int
Average
	Average(dice{}) math.Rat

*/
package main

import (
	"math/big"
	"sort"
)

// Given a series of rolls, determine how to sum them. Use for re-roll, advantage, etc.
type sumFn func(rolls []int) int

type keepFn func(roll int) bool

func defaultSum(rolls []int) int {
	sum := 0
	for _, roll := range rolls {
		sum += roll
	}
	return sum
}

func dropLowest(rolls []int) int {
	rollsCp := make([]int, len(rolls))
	copy(rollsCp, rolls)
	sort.Slice(rollsCp, func(i, j int) bool {
		return rollsCp[i] < rollsCp[j]
	})
	sum := 0
	for i := 1; i < len(rollsCp); i++ {
		sum += rollsCp[i]
	}
	return sum
}

func dropHighest(rolls []int) int {
	rollsCp := make([]int, len(rolls))
	copy(rollsCp, rolls)
	sort.Slice(rollsCp, func(i, j int) bool {
		return rollsCp[i] < rollsCp[j]
	})
	sum := 0
	for i := 0; i < len(rollsCp)-1; i++ {
		sum += rollsCp[i]
	}
	return sum
}

func keepOver2(n int) bool {
	if n > 2 {
		return true
	} else {
		return false
	}
}

// Iterate over pairs, taking the first unless it fails given keepFn
func reroll(keep keepFn, rolls []int) int {
	sum := 0
	for i := 0; i < len(rolls); i += 2 {
		if keep(rolls[i]) {
			sum += rolls[i]
		} else {
			sum += rolls[i+1]
		}
	}
	return sum
}

func rerollOneAndTwo(rolls []int) int {
	return reroll(keepOver2, rolls)
}

// Dice possibilities as a slice
func DS(sides int) []int {
	var r []int
	for i := 1; i <= sides; i++ {
		r = append(r, i)
	}
	return r
}

func Explode(reroll bool, dice ...int) chan []int {
	out := make(chan []int)
	var diceCp []int
	for _, d := range dice {
		if reroll {
			diceCp = append(diceCp, d, d)
		} else {
			diceCp = append(diceCp, d)
		}
	}
	go func() {
		var rolled []int
		rExplode(out, rolled, diceCp)
		close(out)
	}()
	return out
}

func rExplode(out chan []int, rolled, unrolled []int) {
	if len(unrolled) == 1 {
		for _, v := range DS(unrolled[0]) {
			cp := make([]int, len(rolled)+1)
			for idx, r := range rolled {
				cp[idx] = r
			}
			cp[len(rolled)] = v
			out <- cp
		}
	} else {
		for _, v := range DS(unrolled[0]) {
			newRolled := append(rolled, v)
			rExplode(out, newRolled, unrolled[1:])
		}
	}
}

/* Recursive add to discover all possible sums when rolling a group of dice.
out: Channel to return final sums on.
rolled: The values rolled so far.
unrolledDice: Slice where each number represents the sides on a future dice to add. */
func rAdd(out chan int, rolled []int, unrolledDice []int, fn sumFn) {
	//fmt.Println("rolled:", rolled, "unrolledDice:", unrolledDice)
	if len(unrolledDice) == 1 {
		for _, v := range DS(unrolledDice[0]) {
			finalRolls := append(rolled, v)
			out <- fn(finalRolls)
		}
	} else {
		for _, v := range DS(unrolledDice[0]) {
			newRolled := append(rolled, v)
			rAdd(out, newRolled, unrolledDice[1:], fn)
		}
	}
}

/* Figure out all possible sums for a group of dice. */
func DR2(fn sumFn, dices ...int) <-chan int {
	out := make(chan int)
	go func() {
		var rolled []int
		rAdd(out, rolled, dices, fn)
		close(out)
	}()
	return out
}

/* Create a map where keys are possible sums and values are how many ways to achieve that sum
Takes a list of numbers, each one represents how many sides are on the dice. */
func DMap(reroll bool, fn sumFn, dices ...int) map[int]int {
	m := make(map[int]int)
	for v := range Explode(reroll, dices...) {
		s := fn(v)
		m[s] = m[s] + 1
	}
	return m
}

func Average(reroll bool, fn sumFn, dice ...int) float64 {
	total := 0
	count := 0
	for v := range Explode(reroll, dice...) {
		s := fn(v)
		total += s
		count++
	}
	return float64(total) / float64(count)
}

func AverageRat(reroll bool, fn sumFn, dice ...int) *big.Rat {
	total := 0
	count := 0
	for v := range Explode(reroll, dice...) {
		s := fn(v)
		total += s
		count++
	}
	return big.NewRat(int64(total), int64(count))
}
