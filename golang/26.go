package main

import (
	"fmt"
	//	"math/big"
	//	"github.com/fenglyu/projecteuler/golang/common"
)

func CircurLength(n int) int {
	a := 1
	b := n
	t := 0
	hash := make(map[int]int)

	for {
		if hash[a] > 0 {
			break
		}

		hash[a]++
		//		fmt.Println(a, "/", b, "=") //, n)
		a = a % b * 10
		t++
	}
	/*
		fmt.Println(t - hash[a])
	*/

	return t - hash[a]
	// https://www.xarg.org/puzzle/project-euler/problem-26/
}

func main() {

	var i int

	var max int = 0
	var max_p int = 0
	for i = 1; i < 1001; i++ {
		tmp := CircurLength(i)
		if max < tmp {
			max_p = i
			max = tmp
		}
	}

	fmt.Println(max_p)

}
