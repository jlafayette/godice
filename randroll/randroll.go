package randroll

import (
	"math/big"
	"math/rand"
	"time"
)

type sumFn func(rolls []int) int

// R returns a random roll for a dice with the given number of sides.
func R(sides int, r rand.Rand) int {
	return r.Intn(sides) + 1
}

// RD returns a random roll for each dice in a group.
func RD(dice []int, r rand.Rand) []int {
	results := make([]int, len(dice))
	for i, d := range dice {
		results[i] = r.Intn(d) + 1
	}
	return results
}

// RandAverage calculates an average from many random rolls for the given set of dice.
func RandAverage(dice []int, sumfn sumFn, count int, concurrent int) *big.Rat {
	adder := func(n int) chan int {
		c := make(chan int)
		time.Sleep(time.Nanosecond)
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		go func() {
			sum := 0
			for i := 0; i < n; i++ {
				result := RD(dice, *r)
				sum += sumfn(result)
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
	return big.NewRat(int64(finalSum), int64(count))
}
