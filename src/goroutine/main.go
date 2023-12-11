package main

import (
	"fmt"
	"time"
)

const (
	MAXVALUE = 10
)

func main() {
	var sliceV1 = []int{1, 3, 4}
	fmt.Println(cap(sliceV1))

	go func() {
		const n = 48
		fibN := fib(n)
		fmt.Printf("\rFibonacci(%d) = %d", n, fibN)
	}()

	//
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals, MAXVALUE)
	go squarer(squares, naturals)
	go printer(squares)

	//
	go pipelineTest(MAXVALUE)

	// startServ()
	go spinner(100 * time.Millisecond)

	StartCrawl(20)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
