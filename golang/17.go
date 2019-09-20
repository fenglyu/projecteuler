package main

import (
	"fmt"
)

func convert(t int) interface{} {
	stats := map[int]int{}
	n := t
	c := 0
	for {
		if n == 0 {
			break
		}

		stats[c] = n % 10
		n = n / 10
		c++
	}

	return stats
}

func main() {

	ones := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	tens := []string{
		"twenty",
		"thirty",
		"fourty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	fmt.Println(stat)
	fmt.Println(convert(422))
	s := convert(422)

	for i:=2; i>=0 ; i-- {
		fmt.Println(
	}
}
