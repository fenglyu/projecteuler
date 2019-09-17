package main

import (
	"fmt"
)

func prime(n int) bool {
	i := 2
	for {
		if i*i > n {
			break
		}

		if n%i == 0 {
			return false
		}

		i++
	}
	return true
}

func main() {
	i, c := 2, 0

	for {
		if prime(i) {
			c++
			if c >= 10001 {
				fmt.Println(i, c)
				break
			}
		}

		i++
	}

}
