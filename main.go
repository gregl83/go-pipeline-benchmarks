package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	iterations = 1000000
)

func routines() (time.Duration, float32) {
	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func(val int) {
			C(B(A(val)))
			wg.Done()
		}(i)
	}

	wg.Wait()

	elapsed := time.Since(start)

	ms := elapsed / time.Millisecond
	rate := (float32(iterations) / float32(ms)) * 1000

	return ms, rate
}

func channelPipeline() (time.Duration, float32) {
	var wg sync.WaitGroup

	in := make(chan int)
	out := make(chan string)
	defer close(in)
	defer close(out)

	aOut := make(chan int)
	defer close(aOut)

	go func(in chan int, a chan int) {
		for payload := range in {
			go func(p int, o chan int) {
				o <- A(p)
			}(payload, a)
		}
	}(in, aOut)

	bOut := make(chan string)
	defer close(bOut)

	go func(in chan int, b chan string) {
		for payload := range in {
			b <- B(payload)
		}
	}(aOut, bOut)

	go func(in chan string, c chan string) {
		for payload := range in {
			out <- C(payload)
		}
	}(bOut, out)

	var n int

	go func() {
		for range out {
			wg.Done()
		}
	}()

	start := time.Now()
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		in <- n
	}

	wg.Wait()

	elapsed := time.Since(start)

	ms := elapsed / time.Millisecond
	rate := (float32(iterations) / float32(ms)) * 1000

	return ms, rate
}

func printResults(duration time.Duration, rate float32) {
	fmt.Printf("time: %d, rate: %f \r\n", duration, rate)
}

func main() {
	fmt.Println("Welcome to go-pipeline-benchmarks! \r\n")
	fmt.Println("For results run:  go test -bench= \r\n")

	printResults(routines())
	printResults(channelPipeline())
}