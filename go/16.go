package main

import (
	"fmt"
)

func main() {
	c := 1 << 15

	fmt.Println("c: ", c)
	//d := 1 << 1000

	var sum uint64

	sum = 1

	for i := 0; i < 10; i++ {
		sum = sum * 2
	}

	fmt.Println("sum: ", sum)
}
