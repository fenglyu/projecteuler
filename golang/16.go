package main

import (
	"fmt"
	"math"
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
	//result := make([]int, (len(f)+len(g))*2)
	// optimization, take care of memory allocation
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
	a := []int{2, 3, 4, 5, 6}
	b := []int{4, 5, 3, 1}

	result := mul(a, b)
	//	fmt.Println(result)

	sum := 0
	inc := 1
	for i := len(result) - 1; i > 0; i-- {
		sum += result[i] * inc
		inc = inc * 10
	}
	fmt.Println(sum)

	// 2**50
	c := []int{1, 1, 2, 5, 8, 9, 9, 9, 0, 6, 8, 4, 2, 6, 2, 4}
	//d := []int{1, 1, 2, 5, 8, 9, 9, 9, 0, 6, 8, 4, 2, 6, 2, 4}
	res := make([]int, 1000)
	res[999] = 1

	for i := 0; i < 20; i++ {
		res = mul(res, c)
	}

	sum = 0
	for i := len(res) - 1; i >= 0; i-- {
		sum += res[i]
	}
	fmt.Println("sum: ", sum)

	// what the fuck, there is math lib :(
	o := math.Pow(2, 1000)
	fmt.Printf("%.1f", o)
}

/*
result same as python3 result 2**1000

[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 7 1 5 0 8 6 0 7 1 8 6 2 6 7 3 2 0 9 4 8 4 2 5 0 4 9 0 6 0 0 0 1 8 1 0 5 6 1 4 0 4 8 1 1 7 0 5 5 3 3 6 0 7 4 4 3 7 5 0 3 8 8 3 7 0 3 5 1 0 5 1 1 2 4 9 3 6 1 2 2 4 9 3 1 9 8 3 7 8 8 1 5 6 9 5 8
5 8 1 2 7 5 9 4 6 7 2 9 1 7 5 5 3 1 4 6 8 2 5 1 8 7 1 4 5 2 8 5 6 9 2 3 1 4 0 4 3 5 9 8 4 5 7 7 5 7 4 6 9 8 5 7 4 8 0 3 9 3 4 5 6 7 7 7 4 8 2 4 2 3 0 9 8 5 4 2 1 0 7 4 6 0 5 0 6 2 3 7 1 1 4 1 8 7 7 9 5 4 1 8 2 1 5 3 0 4 6 4 7 4 9 8 3 5 8 1 9 4 1 2 6 7 3 9 8 7 6 7 5 5 9 1 6 5 5 4 3 9 4 6 0 7 7 0 6 2 9 1 4 5 7 1 1 9 6 4 7 7 6 8 6 5 4 2 1 6 7 6 6 0 4 2 9 8 3 1 6 5 2 6 2 4 3 8 6 8 3 7 2 0 5 6 6 8 0 6 9 3 7 6]
*/
