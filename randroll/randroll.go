package randroll

import (
	"github.com/jlafayette/godice/mathutil"
	"math/big"
	"math/rand"
	"time"
)

type rollFn func(sides int, r rand.Rand) int

func R(sides int, r rand.Rand) int {
	return r.Intn(sides) + 1
}

func AdvantageR(sides int, r rand.Rand) int {
	return mathutil.Max(R(sides, r), R(sides, r))
}

func DisadvantageR(sides int, r rand.Rand) int {
	return mathutil.Min(R(sides, r), R(sides, r))
}

func RD(dice []int, r rand.Rand) []int {
	results := make([]int, len(dice))
	for i, d := range dice {
		results[i] = r.Intn(d) + 1
	}
	return results
}

func RandAverage(count int, concurrent int, fn rollFn) *big.Rat {
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
	return big.NewRat(int64(finalSum), int64(count))
}
