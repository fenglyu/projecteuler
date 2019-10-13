/*
Let d(n) be defined as the sum of proper common.Divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper common.Divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper common.Divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

package main

import (
	"fmt"

	"github.com/fenglyu/projecteuler/golang/common"
)

func sum1j(ar []int) int {

	s := 0
	for _, v := range ar {
		s += v
	}
	return s
}

func main() {
	sum := 40

	nums := common.Divisor(sum)
	fmt.Println("sum: ", nums)
	// 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;
	fmt.Println("220: ", common.Divisor(220))

	s := 0
	for _, v := range common.Divisor(220) {
		s += v
	}
	fmt.Println("220: ", s)
	// 1, 2, 4, 71 and 142
	fmt.Println("284: ", common.Divisor(284))
	s = 0
	for _, v := range common.Divisor(284) {
		s += v
	}
	fmt.Println("284: ", s)

	fmt.Println("Demonstration completed .........")
	fmt.Println("Problem start ............\n\n")

	ar := make([]int, 100)

	for i := 1; i < 10001; i++ {
		tf := common.Divisor(i)
		f := sum1j(tf)
		ts := common.Divisor(f)
		s := sum1j(ts)
		//fmt.Println(f, s)
		if f != 1 && s != 1 && i == s && i != f {
			fmt.Println(i, f, s)
			ar = append(ar, i, f)
		}
		//	fmt.Println(ar)
	}

	fmt.Println(ar)

	ht := make(map[int]int)
	sum = 0
	for _, v := range ar {
		if _, ok := ht[v]; ok {
			ht[v]++
			continue
		} else {
			sum += v
			ht[v]++
		}
	}

	fmt.Println("Hash Table: ", ht)
	fmt.Println("sum: ", sum)
}
