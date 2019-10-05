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

func solution1() {

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

		res = mul(res, ar)

	}
	//	fmt.Println(res)

	sum := 0
	for _, v := range res {
		sum += v
	}

	fmt.Println(sum)
}

func mul2(a []int, b []int, r []int) { //[]int {

	pos := len(r) - 1

	for i := 0; i < len(r); i++ {
		r[i] = 0
	}

	for i := len(a) - 1; i >= 0; i-- {
		off := len(a) - 1 - i

		for j := len(b) - 1; j >= 0; j-- {
			//			fmt.Println(i, j, pos, off, pos-off, pos-off-1)
			tmp := b[j] * a[i]
			r[pos-off] += tmp % 10
			//	if pos-off-1 > 0 {
			r[pos-off-1] += r[pos-off]/10 + tmp/10
			//	}
			r[pos-off] %= 10
			off++
		}
	}

	//	return r
}

func solution2() {

	r := make([]int, 1000)

	res := make([]int, 200)
	res[199] = 1

	for i := 1; i < 101; i++ {

		ar := []int{0, 0, 0}
		count, c := len(ar)-1, i

		for {
			if c == 0 || count < 0 {
				break
			}

			ar[count] = c % 10

			c = c / 10
			count--
		}

		mul2(res, ar, r)
		//fmt.Println(res[190:], ar, r[990:])
		for j := 0; j < len(res); j++ {
			res[j] = r[len(r)-len(res)+j]
		}
	}

	//	fmt.Println(r)

	sum := 0
	for _, v := range r {
		sum += v
	}

	fmt.Println(sum)
}

func main() {

	solution1()
	solution2()

}
