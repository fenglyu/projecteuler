package common

//import (
//"testing"
//	"fmt"
//)

func NextPermutationV1(c []int, l int, r int) { //bool {
	//func NextPermutation(c interface{}, l int, r int) {

	var i, j int
	for i = r - 1; i >= 0; i-- {
		if c[i] < c[i+1] {
			break
		}
	}

	for j = r; j >= 0; j-- {
		if c[j] > c[i] {
			break
		}
	}

	c[i], c[j] = c[j], c[i]

	i++
	j = r

	for {
		if i >= j {
			break
		}

		c[i], c[j] = c[j], c[i]
		i, j = i+1, j-1
	}
}

func NextPermutationV2(c []int, l int, r int) bool {

	if l >= r {
		return false
	}

	var i, j int
	i = r - 1

	for {
		if c[i] < c[i+1] {
			for j = r; j >= 0; j-- {
				if c[j] > c[i] {
					break
				}
			}
			c[i], c[j] = c[j], c[i]
			ReverseV2(c, i+1, r)
			return true
		}

		i--

		if i == l {
			ReverseV2(c, l, r)
			return false
		}
	}
}

func ReverseV2(c []int, l int, r int) {

	if l == r {
		return
	}

	for {
		if l >= r {
			break
		}
		c[l], c[r] = c[r], c[l]
		l++
		r--
	}
}

// https://github.com/cweill/Permute-Golang/blob/master/permute.go
// http://wordaligned.org/articles/next-permutation
func NextPermutationV3(c []int, first int, last int) bool {

	if first == last {
		return false
	}

	i := first
	i++

	if i == last {
		return false
	}

	i = last - 1

	for {
		ii := i
		i--

		if c[i] < c[ii] {
			j := last - 1
			for {
				if c[i] < c[j] {
					break
				} else {
					j--
				}
			}

			c[i], c[j] = c[j], c[i]
			Reverse(c, ii, last)
			return true
		}

		if i == first {
			Reverse(c, first, last)
			return false
		}
	}
}

// Reverse() is a predefined function in header file algorithm. It is defined as a template in the above mentioned header file. It reverses the order of the elements in the range [first, last) of any container.
//Note: The range used is [first,last), which contains all the elements between first and last, including the element pointed by first but not the element pointed by last.
// https://github.com/gcc-mirror/gcc/blob/gcc-9-branch/libstdc++-v3/include/bits/stl_algo.h#L1150-L1164
func Reverse(c []int, first int, last int) {

	if first == last {
		return
	}

	last--
	for {
		if first >= last {
			break
		}
		c[first], c[last] = c[last], c[first]
		first++
		last--
	}
}
