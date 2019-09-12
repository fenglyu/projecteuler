package main

import (
	"fmt"
	//	set "github.com/deckarep/golang-set"
	//	"sort"
)

func GetPrimes(n int) []int {

	i := 1

	s := make([]int, 21)
	for {

		if n == 1 || n == 2 {
			s.append(s, n)
			return s
		}

		if i*i > n {
			break
		}

		if n%i == 0 {
			s = append(s, i, n/i)
		}

		i += 1
	}

	return s
}

func main() {
	fmt.Println(GetPrimes(10))
}
