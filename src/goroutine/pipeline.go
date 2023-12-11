package main

import (
	"fmt"
	"math"
)

func pipelineTest(maxVal int) {
	naturals := make(chan int)
	squares := make(chan int)

	//
	go func() {
		for x := 0; x < maxVal; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	go func() {
		for {
			fmt.Println(<-squares)
		}
	}()
}

func counter(out chan<- int, count int) {
	for x := 0; x < count; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- int(math.Pow(float64(v), 2))
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
