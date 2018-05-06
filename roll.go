package main

import "fmt"

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
out: int channel to return final sums on.
current: The sum so far.
ns: Slice where each int is the number of sides on a dice. Represents future dice to add.
*/
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

/*  */
func DR2(dices ...int) <-chan int {
	out := make(chan int)
	go func() {
		rAdd(out, 0, dices)
		close(out)
	}()
	return out
}

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
}
