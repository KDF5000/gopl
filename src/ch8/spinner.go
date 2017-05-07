package main

import (
	"fmt"
	"time"
)

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
func fib2(x int) int {
	if x < 2 {
		return x
	}
	a, b := 1, 1
	res := 0
	for i := 3; i <= x; i++ {
		res = a + b
		a, b = b, res
	}
	return res
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	go spinner(100 * time.Millisecond)
	const N = 45
	start := time.Now()
	fibN := fib(N)
	time1 := time.Since(start)
	mid := time.Now()
	fibN2 := fib2(N)
	time2 := time.Since(mid)

	fmt.Printf("\rFibonacci(%d) = %d, Fib2:%d=%d\n", N, fibN, N, fibN2)
	fmt.Printf("Recursion:%v, Loop:%v\n", time1, time2)
}
