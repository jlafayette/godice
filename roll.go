package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func AdvantageR(r rand.Rand) int {
	return max(R(20, r), R(20, r))
}

func DisadvantageR(r rand.Rand) int {
	return min(R(20, r), R(20, r))
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
current: The sum so far.
ns: Slice where each number represents the sides on a future dice to add. */
func rAdd(out chan int, current int, ns []int) {
	if len(ns) == 1 {
		for _, v := range DS(ns[0]) {
			out <- current + v
		}
	} else {
		for _, v := range DS(ns[0]) {
			rAdd(out, current+v, ns[1:])
		}
	}
}

/* Figure out all possible sums for a group of dice. */
func DR2(dices ...int) <-chan int {
	out := make(chan int)
	go func() {
		rAdd(out, 0, dices)
		close(out)
	}()
	return out
}

/* Create a map where keys are possible sums and values are how many ways to achieve that sum
Takes a list of numbers, each one represents how many sides are on the dice. */
func DMap(dices ...int) map[int]int {
	m := make(map[int]int)
	for v := range DR2(dices...) {
		m[v] = m[v] + 1
	}
	return m
}

func main() {
	for _, v := range DS(20) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	for v := range DC(20) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	i := 0
	for v := range DR2(6, 6) {
		fmt.Print(v, " ")
		i++
	}
	fmt.Println("\ni:", i)
	fmt.Println(DMap(6, 6, 6))
	fmt.Println("Testing random rolls:")
	for t := 0; t < 10; t++ {
		TestR(20, 20)
	}
}
