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

	if l == r {
		return false
	}

	var i, j int
	i = r - 1
	//	if l == i {
	//		return false
	//	}

	for {
		if c[i] < c[i+1] {
			for j = r; j >= 0; j-- {
				if c[j] > c[i] {
					break
				}
			}
			c[i], c[j] = c[j], c[i]
			Reverse(c, i+1, r)
			return true
		}

		i--

		if i == l {
			Reverse(c, l, r)
			return false
		}
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
				if !(c[i] < c[j]) {
					j--
					continue
				} else {
					break
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

// reverse() is a predefined function in header file algorithm. It is defined as a template in the above mentioned header file. It reverses the order of the elements in the range [first, last) of any container.
//Note: The range used is [first,last), which contains all the elements between first and last, including the element pointed by first but not the element pointed by last.
func Reverse(c []int, first int, last int) {
	j := last
	last--
	for {
		if first >= j || first == last {
			break
		}
		c[first], c[last] = c[last], c[first]
		first++
		last--
	}
}
