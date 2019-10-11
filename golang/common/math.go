package common

import (
	"fmt"
	"math"
)

func Plus(a []int, b []int, r []int) {
	for i := 0; i < len(r); i++ {
		r[i] = 0
	}

	pos := len(r) - 1

	j := pos
	for {

	}
}

func Multiple(a []int, b []int, r []int) {

	pos := len(r) - 1

	for i := 0; i < len(r); i++ {
		r[i] = 0
	}

	for i := len(a) - 1; i >= 0; i-- {
		off := len(a) - 1 - i

		for j := len(b) - 1; j >= 0; j-- {
			tmp := b[j] * a[i]
			r[pos-off] += tmp % 10
			r[pos-off-1] += r[pos-off]/10 + tmp/10
			r[pos-off] %= 10
			off++
		}
	}
}

func MultipleWithResult(f []int, g []int) []int {

	if len(f) < len(g) {
		t := f
		f = g
		g = t
	}

	result := make([]int, len(f)+len(g)+1)

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

func Divisor(n int) []int {

	nums := make([]int, 500)

	for i := 0; i < len(nums); i++ {
		nums[i] = 0
	}

	i, c := 1, 0

	for {
		if i*i > n {
			break
		}

		tmp := n % i
		if tmp == 0 {
			nums[c] = i
			// only collect 1, ingore the number itself
			if i != 1 && i != n/i {
				c++
				nums[c] = n / i
			}
		}
		i++
		c++
		//fmt.Println("c:  ", c)
	}

	return nums
}

func Fib(n int) int {
	if n <= 2 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

// https://stackoverflow.com/questions/33923/what-is-tail-recursion
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, ...

func TailFib(n int) int {
	if n <= 2 {
		return 1
	}
	return tailRecursiveAux(n, 1, 1)
}

func tailRecursiveAux(n int, iter int, acc int) int {
	//	fmt.Println(n, iter, acc)
	if iter == n {
		return acc
	}
	//return tailRecursiveAux(n, iter++, acc+iter)
	iter++
	return tailRecursiveAux(n, iter, acc+iter)
}

func GoldenFib(n int) float64 {
	f := 1.618034
	m := math.Pow(f, 6) - math.Pow(1-f, 6)
	l := math.Round(m / math.Sqrt(5))
	fmt.Println(f, m, l)
	return l

}

func TailFibLarge(n []int) []int {
	if n <= 2 {
		return 1
	}
	return tailRecursiveAuxLarge(n, 1, 1)
}

func tailRecursiveAuxLarge(n []int, iter int, acc int) int {
	//	fmt.Println(n, iter, acc)
	if iter == n {
		return acc
	}
	//return tailRecursiveAux(n, iter++, acc+iter)
	iter++
	return tailRecursiveAuxLarge(n, iter, acc+iter)
}
