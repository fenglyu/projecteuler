package main

import (
	"fmt"
)

func convert(n int) string {
	//func convert(n int) interface{} {

	ones := []string{
		"",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	tens := []string{
		"",
		"",
		"twenty",
		"thirty",
		"fourty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	stats := make([]int, 4, 5)
	//	n := t
	c := 0
	for {
		if n == 0 {
			break
		}

		stats[c] = n % 10
		n = n / 10
		c++
	}

	//t := s.([]int)

	t := stats
	res := ""

	//fmt.Println(res)
	return res
}

func convert2(n int) []int {
	a := []int{1, 2, 3}
	return a
}

func main() {

	//	fmt.Println(stat)
	//	fmt.Println(convert(422))
	arr := make([]string, 1001, 1050)
	sum := 0
	for i := 1; i < 1001; i++ {
		arr[i-1] = convert(i)
		sum += len(arr[i-1])

		fmt.Println(arr[i-1])
		//fmt.Println(i)
		//sum += len(convert(i))
	}

	fmt.Println(sum)

}
