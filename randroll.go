package dice

import (
	"time"
	"fmt"
	"math/rand"
)

type rollFn func(sides int, r rand.Rand) int

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

func randtesting() {
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
}
