package main

import (
	"fmt"
	//	"math"
	"github.com/fenglyu/projecteuler/golang/common"
)

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

		res = common.MultipleWithResult(res, ar)

	}
	//	fmt.Println(res)

	sum := 0
	for _, v := range res {
		sum += v
	}

	fmt.Println(sum)
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

		common.Multiple(res, ar, r)
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
