package main

import (
	"fmt"
)

func main() {
	var sum int
	sum = fib(1)

	fmt.Printf("sum: %d\n", sum)
	sum = fib(0)
	fmt.Printf("sum: %d\n", sum)

	sum = fib(30)
	fmt.Printf("sum: %d\n", sum)

}

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
