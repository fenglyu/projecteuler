package main

import (
	"fmt"
)

func isPalindromic(m int) bool {

	n := m
	t := 0
	for {
		if n == 0 {
			break
		}
		t = t*10 + n%10
		n = n / 10
	}
	//	fmt.Println(t)
	if t == m {
		return true
	}
	return false
}

func main() {
	//fmt.Println(isPalindromic(9009))

	lp := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			palindrome := i * j
			if isPalindromic(palindrome) {
				//fmt.Println(palindrome)
				//break
				if palindrome > lp {
					lp = palindrome
				}
			}
		}
	}

	fmt.Println(lp)
}
