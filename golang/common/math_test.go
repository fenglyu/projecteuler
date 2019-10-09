package common

import (
	"testing"
)

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
	expected := []int{1, 2, 20, 0, 4, 10, 5, 8}
	sum := Divisor(num)

	if sum != expected {
		t.Errorf("Result: %d, Expected: %d", sum, expected)
	}
}
