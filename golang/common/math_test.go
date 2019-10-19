package common

import (
	//	"fmt"
	"reflect"
	"testing"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestNumbericLength(t *testing.T) {

	nums, expected := []int{0, 9, 8, 9, 8, 7}, 5

	n := NumbericLength(nums)

	if n != expected {
		t.Errorf("Result: %d, Expected: %d", n, expected)
	}
}

func TestPlus(t *testing.T) {

	//	a := []int{9, 4, 4, 5, 6}
	a := []int{1, 9, 9, 7, 6, 5}
	//	b := []int{4, 5, 3, 1}
	b := []int{0, 2, 3, 5, 0}

	expected := []int{0, 2, 0, 2, 1, 1, 5}
	result := make([]int, max(len(a), len(b)+2))
	for i := 0; i < len(result); i++ {
		result[i] = 0
	}

	Plus(a, b, result)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result: %d, Expected: %d", result, expected)
	}
}

func TestPlus2(t *testing.T) {

	sum := 0
	init := []int{0}
	result := make([]int, 20)
	for i := 1234; i <= 4458; i++ {
		sum += i
		f, ar := i, []int{0, 0, 0, 0, 0}
		c := len(ar) - 1
		for {
			if f == 0 {
				break
			}
			ar[c] = f % 10
			f = f / 10
			c--
		}

		Plus(init, ar, result)
		init = make([]int, len(result))
		//fmt.Println(ar, result)
		copy(init, result)
	}

	r := 0
	for i := 0; i < len(result); i++ {
		r = r*10 + int(result[i])
	}
	if sum != r {
		t.Errorf("Result: %d, Expected: %d", r, sum)
	}
}

func TestPlus3(t *testing.T) {

	a := []int{9, 6, 1, 7}
	b := []int{1, 8, 1, 3}
	expected := []int{1, 1, 4, 3, 0}

	result := make([]int, max(len(a), len(b)+2))
	for i := 0; i < len(result); i++ {
		result[i] = 0
	}

	Plus(a, b, result)

	n := len(result) - NumbericLength(result)
	if !reflect.DeepEqual(result[n:], expected) {
		t.Errorf("Result: %d, Expected: %d", result, expected)
	}
}

/*
func TestPlusV2(t *testing.T) {

	sum := 0
	init := []int{0}
	result := make([]int, 20)
	for i := 1234; i <= 4458; i++ {
		sum += i
		f, ar := i, []int{0, 0, 0, 0, 0}
		c := len(ar) - 1
		for {
			if f == 0 {
				break
			}
			ar[c] = f % 10
			f = f / 10
			c--
		}
		PlusV2(init, ar, result)
		init = make([]int, len(result))
		//fmt.Println(ar, result)
		copy(init, result)
	}

	r := 0
	for i := 0; i < len(result); i++ {
		r = r*10 + int(result[i])
	}
	if sum != r {
		t.Errorf("Result: %d, Expected: %d", r, sum)
	}
}
*/

func TestMultipleWithResultSmall(t *testing.T) {
	a := []int{2, 3, 4, 5, 6}
	b := []int{4, 5, 3, 1}
	expected := 106279136

	result := MultipleWithResult(a, b)

	sum := 0
	inc := 1
	for i := len(result) - 1; i > 0; i-- {
		sum += result[i] * inc
		inc = inc * 10
	}
	if sum != expected {
		t.Errorf("Result: %d, Expected: %d", sum, expected)
	}
}

func TestMultiple(t *testing.T) {

	r := make([]int, 1000)
	res := make([]int, 200)
	res[199] = 1
	expected := 648

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
		Multiple(res, ar, r)
		for j := 0; j < len(res); j++ {
			res[j] = r[len(r)-len(res)+j]
		}
	}

	sum := 0
	for _, v := range r {
		sum += v
	}

	if sum != expected {
		t.Errorf("Result: %d, Expected: %d", sum, expected)
	}
}

func TestDivisor(t *testing.T) {
	num := 220
	expected := []int{1, 2, 110, 0, 4, 55, 5, 44, 0, 0, 0, 0, 10, 22, 11, 20, 0}
	sum := Divisor(num)

	if !reflect.DeepEqual(sum[:17], expected) {
		t.Errorf("Result: %d, Expected: %d", sum, expected)
	}
}

func TestDivideV1(t *testing.T) {
	a, b := 32, 5
	eq, er := 6, 2

	q, r := DivideV1(a, b)

	if q != eq || r != er {
		t.Errorf("Result: %d, Expected: %d", q, eq)
		t.Errorf("Result: %d, Expected: %d", r, er)
	}
}

// https://www.wikiwand.com/zh-cn/带余除法#/cite_ref-hw_1-2
func TestQuickDivision(t *testing.T) {

	// Test 1
	a, b := 237, 9
	eq, er := 26, 3
	q, r := QuickDivision(a, b)
	if q != eq || r != er {
		t.Errorf("Result: %d, Expected: %d", q, eq)
		t.Errorf("Result: %d, Expected: %d", r, er)
	}
	// Test 2
	a, b = 234324342342, 29
	eq, er = 8080149735, 27
	q, r = QuickDivision(a, b)
	if q != eq || r != er {
		t.Errorf("Result: %d, Expected: %d", q, eq)
		t.Errorf("Result: %d, Expected: %d", r, er)
	}
}

func TestMinusV1(t *testing.T) {

	sum := 5000
	init := []int{1}
	a := []int{5, 0, 0, 0}
	r := make([]int, 20)
	for i := 0; i <= 4000; i++ {
		sum -= 1
		//fmt.Println(a, init, r)
		PostiveSub(a, init, r)
		//		fmt.Println(a, init, r)
		na := make([]int, len(a))
		copy(na, r[(len(r)-len(a)):])
		a = na
		//		fmt.Println(a, init, r)
	}

	re := 0
	for i := 0; i < len(r); i++ {
		re = re*10 + int(r[i])
	}

	//	fmt.Println(re, sum)

	if sum != re {
		t.Errorf("Result: %d, Expected: %d", re, sum)
	}
}
