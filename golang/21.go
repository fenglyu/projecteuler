/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

package main

import (
	"fmt"
)

func divisor(n int) []int {

	nums := make([]int, 500)

	for i := 0; i < len(nums); i++ {
		nums[i] = 0
	}

	i, c := 1, 0

	for {
		if i*i > n {
			break
		}

		tmp := n % i
		if tmp == 0 {
			nums[c] = i
			// only collect 1, ingore the number itself
			if i != 1 && i != n/i {
				c++
				nums[c] = n / i
			}
		}
		i++
		c++
		//fmt.Println("c:  ", c)
	}

	return nums
}

func sum1j(ar []int) int {

	s := 0
	for _, v := range ar {
		s += v
	}
	return s
}

func main() {
	//	sum := 40
	//
	//	nums := divisor(sum)
	//	fmt.Println("sum: ", nums)
	//	// 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;
	//	fmt.Println("220: ", divisor(220))
	//
	//	s := 0
	//	for _, v := range divisor(220) {
	//		s += v
	//	}
	//	fmt.Println("220: ", s)
	//	// 1, 2, 4, 71 and 142
	//	fmt.Println("284: ", divisor(284))
	//s = 0
	//for _, v := range divisor(284) {
	//	s += v
	//}
	//fmt.Println("284: ", s)

	btk := make([]int, 10001)
	ar := make([]int, 100)

	for i := 1; i < 10001; i++ {
		tf := divisor(i)
		f := sum1j(tf)
		ts := divisor(f)
		s := sum1j(ts)
		//fmt.Println(i, tf, ts, s)
		if f == s {
			ar = append(ar, f, s)
		}
		//	fmt.Println(ar)
	}

	fmt.Println(ar)
}
