package main

import (
	"fmt"
)

func prime(n int) bool {
	i := 2
	for {
		if i*i > n {
			break
		}

		if n%i == 0 {
			return false
		}

		i++
	}
	return true
}

func main() {

	sum := 0
	for i := 2; i < 2000000; i++ {
		if prime(i) {
			sum += i
		}
	}

	fmt.Printf("sum: %d\n", sum)
}
