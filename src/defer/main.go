package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

func main() {
	bigSlowOperation()

	f(3)
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()

	time.Sleep(300 * time.Millisecond)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func MustCompile(expr string) *regexp.Regexp {
	re, err := regexp.Compile(expr)
	if err != nil {
		panic(err)
	}
	return re
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
