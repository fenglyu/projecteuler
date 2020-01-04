package main

import (
	"fmt"
	"math"
)

// the top boundry is 9**5 * 6 = 354294
var MAX_N = 354294

func isNumber(n int) bool {
	var sum int = 0
	x := n

	for x != 0 {
		sum += int(math.Pow(float64(x%10), float64(5)))
		x /= 10
	}

	return sum == n
}

func main() {
	sum := 0
	for i := 2; i <= MAX_N; i++ {
		if isNumber(i) {
			sum += i
		}
	}

	fmt.Println("sum %d\n", sum)
}
