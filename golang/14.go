package main

import (
	"fmt"
)

func main() {

	stats := [1000001]int{}

	for i := 1000000; i > 1; i-- {

		count := 0
		n := i
		for {
			if n == 1 {
				break
			}

			if n%2 == 0 {
				n = n / 2
			} else {
				n = 3*n + 1
			}

			count++
		}

		stats[i] = count
	}

	//	fmt.Printf("sum:", stats)

	max := 0
	for i, v := range stats {
		if v > stats[max] {
			max = i
		}
	}

	fmt.Println("max: ", max, "result: ", stats[max])
}
