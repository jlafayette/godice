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
	"math/rand"
	"sort"
	"math/big"
)

type rollFn func(sides int, r rand.Rand) int

// Given a series of rolls, determine how to sum them. Use for re-roll, advantage, etc.
type sumFn func(rolls []int) int

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

func rerollBelow2(rolls []int) int {
	if len(rolls) < 2 {
		return rolls[0]
	}
	if rolls[len(rolls)-2] <= 2 {
		return rolls[len(rolls)-1]
	}
	return rolls[len(rolls)-2]
}

// Dice possibilities as a slice
func DS(sides int) []int {
	var r []int
	for i := 1; i <= sides; i++ {
		r = append(r, i)
	}
	return r
}

// Dice possibilities as a channel
func DC(sides int) <-chan int {
	var out = make(chan int)
	go func() {
		for i := 1; i <= sides; i++ {
			out <- i
		}
		close(out)
	}()
	return out
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
			out <- append(rolled, v)
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

func Average(fn sumFn, dice ...int) float64 {
	total := 0
	count := 0
	for v := range DR2(fn, dice...) {
		total += v
		count++
	}
	return float64(total) / float64(count)
}

func AverageRat(fn sumFn, dice ...int) *big.Rat {
	total := 0
	count := 0
	for v := range DR2(fn, dice...) {
		total += v
		count ++
	}
	return big.NewRat(int64(total), int64(count))
}
