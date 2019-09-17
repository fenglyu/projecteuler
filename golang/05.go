package main

import (
	"fmt"
	//	set "github.com/deckarep/golang-set"
	//	"sort"
)

func GetPrimes(n int) []int {

	s := make([]int, n+1)
	for i := 2; i*i <= n; i++ {
		for {
			if n%i == 0 {
				s[i] += 1
				n = int(n / i)
			} else {
				break
			}
		}
	}

	if n > 1 {
		s[n] += 1
	}

	//	fmt.Println(s)
	return s
}

func main() {
	//fmt.Println(GetPrimes(10))

	s := make([]int, 21)

	for i := 1; i <= 20; i++ {
		//		p := GetPrimes(i)
		//		fmt.Println(p)
		for j, x := range GetPrimes(i) {
			if x > s[j] {
				s[j] = x
			}
		}
	}

	sum := 1

	for i, v := range s {
		for {
			if v <= 0 {
				break
			}

			sum = sum * i
			v -= 1
		}
	}
	//fmt.Println(s)
	fmt.Println(sum)
}
