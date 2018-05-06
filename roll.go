package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
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

func TestAverage(fn sumFn, sides ...int) {
	total := 0
	count := 0
	for v := range DR2(fn, sides...) {
		total += v
		count++
	}
	average := float64(total) / float64(count)
	fmt.Println("For", sides, "average is", average)
}

func R(sides int, r rand.Rand) int {
	return r.Intn(sides) + 1
}

func TestR(sides int, count int) {

	time.Sleep(time.Nanosecond)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	m := make(map[int]int)
	for i := 0; i < count; i++ {
		result := R(sides, *r)
		m[result] = m[result] + 1
		fmt.Printf("%2d ", result)
	}
	fmt.Println(m)
}

func TestRandAverage(count int, concurrent int, fn rollFn) {
	adder := func(n int) chan int {
		c := make(chan int)
		time.Sleep(time.Nanosecond)
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		go func() {
			sum := 0
			for i := 0; i < n; i++ {
				result := fn(20, *r)
				sum += result
			}
			c <- sum
			close(c)
		}()
		return c
	}
	chunkSize := count / concurrent
	remainder := count % concurrent
	outChans := make([]chan int, concurrent+1)
	for i := 0; i < concurrent+1; i++ {
		n := chunkSize
		if i == concurrent {
			n = remainder
		}
		//println("chunkSize", n)
		outChans[i] = adder(n)
	}
	finalSum := 0
	for _, c := range outChans {
		for v := range c {
			//fmt.Println("Adding", v)
			finalSum += v
		}
	}
	fmt.Println("Average with count", count, ":", float64(finalSum)/float64(count))
}

func AdvantageR(sides int, r rand.Rand) int {
	return Max(R(sides, r), R(sides, r))
}

func DisadvantageR(sides int, r rand.Rand) int {
	return Min(R(sides, r), R(sides, r))
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
func DMap(fn sumFn, dices ...int) map[int]int {
	m := make(map[int]int)
	for v := range DR2(fn, dices...) {
		m[v] = m[v] + 1
	}
	return m
}

func main() {
	fmt.Println(DMap(defaultSum, 6, 6, 6))
	fmt.Println("Testing random rolls:")
	for t := 0; t < 10; t++ {
		TestR(20, 20)
	}
	start := time.Now()
	TestRandAverage(1000000, 20, R)
	TestRandAverage(1000000, 20, AdvantageR)
	TestRandAverage(1000000, 20, DisadvantageR)
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", elapsed)

	fmt.Println("Averages from frequency")
	TestAverage(defaultSum, 20)
	TestAverage(defaultSum, 12)
	TestAverage(defaultSum, 10)
	TestAverage(defaultSum, 8)
	TestAverage(defaultSum, 6)
	TestAverage(defaultSum, 4)
	TestAverage(defaultSum, 6, 6)
	TestAverage(defaultSum, 6, 6, 6)
	fmt.Print("(Disadvantage) ")
	TestAverage(dropHighest, 20, 20)
	fmt.Print("(Advantage) ")
	TestAverage(dropLowest, 20, 20)
	fmt.Print("(Drop lowest) ")
	TestAverage(dropLowest, 6, 6, 6, 6)
	fmt.Print("(Default) ")
	TestAverage(defaultSum, 6, 6, 6)

	fmt.Print("(Reroll 1&2) ")
	TestAverage(rerollBelow2, 12, 12)
	fmt.Print("(Reroll 1&2) ")
	TestAverage(rerollBelow2, 6, 6)
}
