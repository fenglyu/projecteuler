package main

import (
	"fmt"
	//	"math/big"
	//	"github.com/fenglyu/projecteuler/golang/common"
)

func CircurLength(n int) int {
	a := 1
	b := n
	c := 0
	t := 0
	hash := make(map[int]int)

	for {
		if c >= 100 || hash[a] > 0 {
			break
		}

		hash[a]++
		//		fmt.Println(a, "/", b, "=") //, n)
		a = a % b * 10
		t++
		c++
	}
	/*
		fmt.Println(t, c)
		fmt.Println(hash)
		fmt.Println(t - hash[a])
	*/

	return t - hash[a]
	// https://www.xarg.org/puzzle/project-euler/problem-26/
}

func main() {

	/*
		fmt.Print("   x")
		for mode := big.ToNearestEven; mode <= big.ToPositiveInf; mode++ {
			fmt.Printf("  %s", mode)
		}
		f64 := 1.6180339887499

		for mode := big.ToNearestEven; mode <= big.ToPositiveInf; mode++ {
			f := new(big.Float).SetPrec(2).SetMode(mode).SetFloat64(f64)
			fmt.Printf("  %*g", len(mode.String()), f)
		}
	*/

	//	fmt.Println("   ")
	//	a := big.NewRat(1, 5)
	//	fmt.Println(a)

	//	fmt.Println(a.FloatString(100))

	//fmt.Println(a.RatString())
	/*
		var i int64
		for i = 2; i < 1001; i++ {
			b := big.NewRat(1, i)
			fmt.Println(b.FloatString(100))
			//fmt.Println(i, b.FloatString(100))
		}
	*/

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
