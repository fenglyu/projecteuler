package common

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	var fib int64 = 0
	var gold float64 = 0

	var i int64
	for i = 1; i <= 20; i++ {
		gold = GoldenFib(int64(i))
		fib = Fib(i)
		if !floatcompare(float64(fib), gold) {
			t.Errorf("Result: %d, Expected: %f", fib, gold)
		}
	}
}

//
func TestTailFib(t *testing.T) {
	var fib int64 = 0
	var gold float64 = 0

	var i int64
	for i = 1; i <= 50; i++ {
		fib = TailFib(i, 0, 1)
		gold = GoldenFib(i)
		// fmt.Printf("[TestTailFib] Result: %d, Expected: %f\n", fib, gold)
		if !floatcompare(float64(fib), gold) {
			t.Errorf("The %d number is Result: %d, Expected: %f", i, fib, gold)
		}
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

func TestTailFibArray(t *testing.T) {
	expected := []int{1, 2, 5, 8, 6, 2, 6, 9, 0, 2, 5}
	i := 50
	fib := TailFibArray(NumToArray(int(i)), []int{0}, []int{1})
	al, bl := len(fib)-NumbericLength(fib), len(expected)-NumbericLength(expected)
	if !reflect.DeepEqual(fib[al:], expected[bl:]) {
		t.Errorf("Result: %d, Expected: %d", fib[al:], expected[bl:])
	}
}

func BenchGoldenFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoldenFib(200)
	}
}

func TestFibBigInt(t *testing.T) {
	var i int64 = 10
	//	for i = 1; i <= 5; i++ {
	bi := big.NewInt(i)
	gold := FibBigInt(bi)

	fib := TailFibArray(NumToArray(int(i)), []int{0}, []int{1})
	fi := len(fib) - NumbericLength(fib)
	f := big.NewInt(ArrayToNum(fib[fi:]))

	//fmt.Println(f, gold)
	if gold.Cmp(f) != 0 {
		t.Errorf("The %d number is Result: %d, Expected: %s", i, f, gold)
	}
}

func TestGoldenFibBig(t *testing.T) {
	var i int64 = 10
	bi := big.NewInt(i)
	gold := GoldenFibBig(bi)
	fmt.Println(gold)
}
