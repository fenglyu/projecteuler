package common

import (
	"fmt"
	"math"
	"math/big"
	//	"reflect"
)

func Fib(n int64) int64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

// https://stackoverflow.com/questions/33923/what-is-tail-recursion
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, ...

func TailFib(n int64, a int64, b int64) int64 {
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	return TailFib(n-1, b, a+b)
}

func TailFibArray(n []int, a []int, b []int) (r []int) {
	//	fmt.Println(n, a, b)
	i := len(n) - NumbericLength(n)
	idx := ArrayToNum(n[i:])
	if idx == 0 {
		return a
	}
	if idx == 1 {
		return b
	}

	r1 := make([]int, len(n))
	PostiveSub(n, []int{1}, r1)
	r = make([]int, len(n)+1)
	Plus(a, b, r)
	//	fmt.Println(n, a, b, r)
	return TailFibArray(r1, b, r)
}

func GoldenFib(n int64) float64 {
	if n <= 2 {
		return 1
	}

	var c float64 = 1
	var t int64 = 2

	for {
		if t >= n {
			break
		}
		c = math.Round(float64(c) * 1.6180339887499)
		t++
	}
	return c
}

func ArrayToNum(nums []int) int64 {
	var s int64 = 0
	for i := 0; i < len(nums); i++ {
		s = s*10 + int64(nums[i])
	}
	return s
}

func NumToArray(num int) []int {

	f, ar := num, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	c := len(ar) - 1
	for {
		if f == 0 {
			break
		}
		ar[c] = f % 10
		f = f / 10
		c--
	}
	return ar
}

func FibBigInt(n *big.Int) *big.Int {

	a := big.NewInt(1)
	b := big.NewInt(1)
	two := big.NewInt(2)

	if n.Cmp(two) <= 0 {
		return a
	}

	i := big.NewInt(1)
	for {
		if n.Cmp(i) <= 0 {
			break
		}

		a.Add(a, b)
		a, b = b, a

		i.Add(i, big.NewInt(1))
		//		fmt.Println(a, b)
	}
	return a
}

func GoldenFibBig(n *big.Int) *big.Float {

	a := big.NewFloat(1)
	two := big.NewInt(2)

	if n.Cmp(two) <= 0 {
		return a
	}

	c := big.NewFloat(1)
	t := big.NewInt(1)

	goldenRate := big.NewFloat(1.6180339887499)

	for {
		if t.Cmp(n) >= 0 {
			break
		}
		c.Mul(c, goldenRate)
		//fmt.Println(c)
		f := new(big.Float).SetPrec(2).SetMode(big.ToNearestAway).Set(c)
		fmt.Println(c, f)
		c = f

		t.Add(t, big.NewInt(1))
	}

	return c
}
