package main

import (
	"fmt"
)

func main() {
	c, sum := 1, 0

	for {
		if c >= 1000 {
			break
		}

		if c%3 == 0 || c%5 == 0 {
			fmt.Printf("c: %d\n", c)
			sum += c
		}

		c++
	}

	fmt.Printf("sum: %d\n", sum)

}
