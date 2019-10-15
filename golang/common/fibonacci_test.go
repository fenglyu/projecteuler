package common

import (
	"math"
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
const float64EqualityThreshold = 1e-9

func floatcompare(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold*(math.Abs(a)+math.Abs(b))
}

func TestGoldenFib(t *testing.T) {

	//var num int64 = 10
	//	var expected float64 = 55
	var num int64 = 1
	var expected float64 = 1

	fib := GoldenFib(num)

	if !floatcompare(float64(fib), float64(expected)) {
		t.Errorf("Result: %f, Expected: %f", fib, expected)
	}
}

func TestTailFibLarge(t *testing.T) {
	nums, expected := []int{0, 1, 0}, []int{0, 5, 5}
	fib := TailFibLarge(nums)
	//fmt.Println(fib, expected)
	al, bl := len(fib)-NumbericLength(fib), len(expected)-NumbericLength(expected)
	if !reflect.DeepEqual(fib[al:], expected[bl:]) {
		t.Errorf("Result: %d, Expected: %d", fib, expected)
	}
}

func TestGoldenFibV2(t *testing.T) {

	var gold float64
	var fib []int

	for i := 1; i <= 101; i++ {
		gold = GoldenFib(int64(i))
		fib = TailFibLarge(NumToArray(i))
		fi := len(fib) - NumbericLength(fib)
		f := ArrayToNum(fib[fi:])

		if !floatcompare(f, gold) {
			t.Errorf("Result: %f, Expected: %f", f, gold)
		}
	}
}

func BenchGoldenFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoldenFib(200)
	}
}
