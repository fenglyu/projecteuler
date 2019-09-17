package main

import (
	"fmt"
)

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

func main() {

	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			if int(2*i*j-2000*i-2000*j+1000000) == 0 {
				fmt.Println(i, j, 1000-i-j)
				fmt.Println(i * j * (1000 - i - j))
			}
		}
	}

	//	fmt.Printf("sum: %d\n", sum)
}
