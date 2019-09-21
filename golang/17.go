package main

import (
	"fmt"
	//	"strings"
)

func convert(z int) string {
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
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	stats := make([]int, 4, 5)
	n := z
	c := 0
	for {
		if n == 0 {
			break
		}

		stats[c] = n % 10
		n = n / 10
		c++
	}

	t := stats
	res := ""

	//t := s.([]int)

	for i := 3; i >= 0; i-- {

		if i == 3 && t[i] != 0 {
			res += fmt.Sprintf("%s thousand", ones[t[i]])
		} else if i == 2 && t[i] != 0 {
			res += fmt.Sprintf("%s hundred", ones[t[i]])
			if t[i-1] != 0 || t[i-2] != 0 {
				res += fmt.Sprint(" and ")
			}
		} else if i == 1 {
			if t[i] == 1 {
				temp := t[i]*10 + t[i-1]
				res += fmt.Sprintf(ones[temp])
			} else {
				res += fmt.Sprintf("%s", tens[t[i]])
				if t[i-1] != 0 && t[i] != 0 {
					res += fmt.Sprintf("-")
				}
			}
		} else if i == 0 {
			if t[i+1] != 1 {
				res += fmt.Sprint(ones[t[i]])
			}
		}
	}

	//fmt.Println(res)
	//fmt.Println(t)

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
	//for i := 1; i < 10; i++ {
	//for i := 10; i < 20; i++ {
	//	for i := 20; i < 100; i++ {
	for i := 100; i < 1000; i++ {
		arr[i-1] = convert(i)
		//	s := strings.ReplaceAll(arr[i-1], " ", "")
		//	s = strings.ReplaceAll(s, "-", "")
		//		fmt.Println(s)
		//sum += len(s)
		for _, v := range arr[i-1] {
			if v != ' ' && v != '-' {
				sum += 1
			}
		}

		//fmt.Println(arr[i-1])
		//fmt.Println(i)
		//sum += len(convert(i))
	}

	fmt.Println(sum)

}
