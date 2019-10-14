package common

import (
	//	"math"
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	num, expected := 10, 55
	fib := Fib(num)
	if fib != expected {
		t.Errorf("Result: %d, Expected: %d", fib, expected)
	}
}

//
func TestTailFib(t *testing.T) {
	num, expected := 10, 55
	fib := TailFib(num)
	if fib != expected {
		//if !floatcompare(float64(fib), float64(expected)) {
		t.Errorf("Result: %d, Expected: %d", fib, expected)
	}
}

//const EPSILON = 1e-9
/*
const float64EqualityThreshold = 1e-9

func floatcompare(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold*(math.Abs(a)+math.Abs(b))
}

func TestGoldenFib(t *testing.T) {

	num, expected := 10, 55
	fib := GoldenFib(num)

	//	fmt.Println(fib, expected)
	if !floatcompare(float64(fib), float64(expected)) {
		t.Errorf("Result: %d, Expected: %d", int(fib), expected)
	}
}
*/

func TestTailFibLarge(t *testing.T) {
	nums, expected := []int{0, 1, 0}, []int{0, 5, 5}
	fib := TailFibLarge(nums)
	//fmt.Println(fib, expected)
	al, bl := len(fib)-NumbericLength(fib), len(expected)-NumbericLength(expected)
	if !reflect.DeepEqual(fib[al:], expected[bl:]) {
		t.Errorf("Result: %d, Expected: %d", fib, expected)
	}
}
