package main

import (
	"fmt"
	"github.com/fenglyu/projecteuler/golang/common"
)

const (
	Deficient = -1
	Perfect   = 0
	Abundant  = 1
)

func NumType(n int) int {

	//nums := divisor(n)
	nums := common.Divisor(n)
	//fmt.Println("c:  ", nums)
	s := 0
	for _, v := range nums {
		s += v
	}

	flag := Perfect
	if s < n {
		flag = Deficient
	} else if s > n {
		flag = Abundant
	}

	//	fmt.Println(s, n, flag)
	return flag
}

func main() {
	//	fmt.Println(Perfect, Deficient, Abundant)
	//fmt.Printf("28 is %v\n", NumType(28))
	//fmt.Printf("12 is %v\n", NumType(12))
	limit := 28123
	sts := make([]int, limit+1)
	for i := 0; i < limit+1; i++ {
		sts[i] = 0
	}

	// collect result
	//	result, rdx := make([]int, limit), 0
	result := make([]int, limit+1)
	for i := 0; i < limit+1; i++ {
		result[i] = 0
	}

	abs, c := make([]int, 10000), 0

	for i := 1; i <= limit; i++ {
		if NumType(i) > 0 {
			//fmt.Println(i)
			abs[c] = i
			sts[i]++
			c++
		}
		//fmt.Println(result)
	}

	for i := 1; i <= limit; i++ {
		var j int
		for j = 0; j <= c; j++ {
			if abs[j] < i {
				div := i - abs[j]
				if sts[div] > 0 {
					result[i]++
					break
				}
			} else {
				break
			}
		}

		// if there is no two abundants found in abs, this number cannot be written as the sum of two abundants numbers.
		if j > c {
			result[i]++
		}
	}

	sum := 0
	for i := 1; i <= limit; i++ {
		if result[i] == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

}
