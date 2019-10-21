package main

import (
	"fmt"
	"math/big"
)

type Float struct {
	integer int64
	point   int64
}

func div(a int, b int) {
	fmt.Println(a, b)
}

type pair struct {
	a int
	b int
}

func main() {

	//	maps := make(map[pair]int)
	var a, b int64
	for a = -1000; a < 1001; a++ {
		for b = -1000; b < 1001; b++ {
			var i int64 = 0
			n := big.NewInt(i)
			c := big.NewInt(1).Mul(n, n)
			g := big.NewInt(a)
			e := big.NewInt(b)
			d := big.NewInt(1).Mul(g, n)
			f := big.NewInt(0).Add(c, d)
			f.Add(f, e)
			for {
				//fmt.Println(a, b)
				if i > 1000 {
					break
				}
				if f.ProbablyPrime(20) {
					fmt.Println(a, b, i, f)
				}
			}

		}
	}
}
