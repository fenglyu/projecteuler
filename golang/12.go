package main

import (
	"fmt"
)

func divisor(n int) int {

	num, i := 0, 1
	for {
		if i*i > n {
			break
		}

		if n%i == 0 {
			num += 2
		}
		i += 1
	}

	if i*i == n {
		num -= 1
	}

	return num
}

func main() {

	i := 1
	sum := 0
	for {
		sum = divisor(i * (i + 1) / 2)
		if sum >= 500 {

			fmt.Printf("sum: %d, num: %d\n", sum, i*(i+1)/2)
			break
		}
		i++
	}
}
