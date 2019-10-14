package common

import (
	"reflect"
)

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

func TailFibLarge(n []int) []int {
	if reflect.DeepEqual(n, []int{1}) || reflect.DeepEqual(n, []int{2}) {
		return []int{1}
	}
	return tailRecursiveAuxLarge(n, []int{1}, []int{1})
}

func tailRecursiveAuxLarge(n []int, iter []int, acc []int) []int {
	//	fmt.Println(n, iter, acc)
	//	if iter == n {

	i, c := len(iter)-NumbericLength(iter), len(n)-NumbericLength(n)

	if reflect.DeepEqual(iter[i:], n[c:]) {
		return acc
	}

	r1 := make([]int, len(iter)+1)
	Plus(iter, []int{1}, r1)
	iter = r1

	r2 := make([]int, len(acc)+1)
	Plus(acc, iter, r2)

	return tailRecursiveAuxLarge(n, iter, r2)
	//return tailRecursiveAux(n, iter++, acc+iter)
}
