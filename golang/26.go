package main

import (
	"fmt"
	"math/big"
	//	"github.com/fenglyu/projecteuler/golang/common"
)

func main() {

	fmt.Print("   x")
	for mode := big.ToNearestEven; mode <= big.ToPositiveInf; mode++ {
		fmt.Printf("  %s", mode)
	}
	f64 := 1.6180339887499

	for mode := big.ToNearestEven; mode <= big.ToPositiveInf; mode++ {
		f := new(big.Float).SetPrec(2).SetMode(mode).SetFloat64(f64)
		fmt.Printf("  %*g", len(mode.String()), f)
	}

	fmt.Println("   ")
	a := big.NewRat(1, 5)
	fmt.Println(a)

	fmt.Println(a.FloatString(100))

	fmt.Println(a.RatString())

	var i int64
	for i = 2; i < 1001; i++ {
		b := big.NewRat(1, i)
		fmt.Println(b.FloatString(100))
	}
}
