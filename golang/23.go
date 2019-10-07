package main

import (
	"fmt"
)

const (
	Deficient = -1
	Perfect   = 0
	Abundant  = 1
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
	}

	return nums
}

func NumType(n int) int {

	nums := divisor(n)
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
