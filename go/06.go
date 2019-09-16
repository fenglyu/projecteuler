package main

import (
	"fmt"
)

func main() {
	c := 100

	sum1 := 0
	sum2 := 0
	for i := 1; i <= c; i++ {
		sum1 += i * i
		sum2 += i
	}

	fmt.Println(sum2*sum2 - sum1)
}
