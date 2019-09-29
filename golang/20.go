package main

import (
	"fmt"
	//	"math"
)

func mul(a interface{}, b interface{}) []int {

	f := a.([]int)
	g := b.([]int)

	if len(f) < len(g) {
		t := f
		f = g
		g = t
	}
	//fmt.Println("f: ", f, "g: ", g)
	result := make([]int, len(f)+len(g))

	for j := len(g) - 1; j >= 0; j-- {
		//initially pos (index of result ) has the same relative postion as j in array
		pos := len(result) - (len(g) - j)
		for i := len(f) - 1; i >= 0; i-- {
			temp := g[j] * f[i]
			result[pos] += temp % 10
			result[pos-1] += temp / 10
			pos--
		}
		//	fmt.Println(result)
	}

	for i := len(result) - 1; i > 0; i-- {
		temp := result[i]
		result[i] = temp % 10
		result[i-1] += temp / 10
	}

	return result
}

func main() {

	//	sum := 0
	res := make([]int, 1000)
	res[999] = 1
	for i := 1; i < 101; i++ {

		count, c := 2, i
		ar := []int{0, 0, 0}

		for {
			if c == 0 {
				break
			}

			ar[count] = c % 10

			c = c / 10
			count--
		}
		//	fmt.Println(c, ar)

		res = mul(res, ar)

	}
	fmt.Println(res)
	sum := 0
	for _, v := range res {
		sum += v
	}
	fmt.Println(sum)

}
