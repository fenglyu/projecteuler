package main

import (
	"fmt"

	"github.com/fenglyu/projecteuler/golang/common"
)

func main() {
	// problem 25
	result := make([]int, 1001)
	first, second := []int{1}, []int{1}

	i := 3

	for {

		common.Plus(first, second, result)
		first = second
		tmp := make([]int, len(result))
		copy(tmp, result)
		second = tmp

		ll := common.NumbericLength(result)
		if ll >= 1000 {
			idx := len(result) - ll
			fmt.Println(len(result[idx:]), result[idx:])
			break
		}

		i++
	}

	//	fmt.Println(result)
	fmt.Println("index of the first term in the Fibonacci sequence to contain 1000 digits", i)

}
