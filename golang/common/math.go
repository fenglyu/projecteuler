package common

import (
	"math"
)

// used for numberic arrays
func NumbericLength(n []int) int {
	i := 0
	for i = 0; i < len(n); i++ {
		if n[i] > 0 {
			break
		}
	}
	return len(n) - i
}

func Plus(a []int, b []int, r []int) {

	for i := 0; i < len(r); i++ {
		r[i] = 0
	}

	al, bl := NumbericLength(a), NumbericLength(b)

	j := len(r) - 1
	c := 0

	pos, pa, pb := 0, 0, 0
	for {
		pos = len(r) - 1 - j

		if c >= al && c >= bl {
			break
		}

		pos = len(r) - 1 - j
		pa = len(a) - 1 - pos
		pb = len(b) - 1 - pos

		if pa >= 0 && pb >= 0 {
			tmp := a[pa] + b[pb]
			r[j] += tmp % 10
			r[j-1] += r[j]/10 + tmp/10
			r[j] %= 10
		} else if pa < 0 {
			r[j] += b[pb]
			r[j-1] += r[j] / 10
			r[j] %= 10
		} else if pb < 0 {
			r[j] += a[pa]
			r[j-1] += r[j] / 10
			r[j] %= 10
		}

		c++
		j--
	}
}

func PlusV2(a []int, b []int, r []int) {
	pos := len(a)
	carry := 0

	for {
		pos--
		if pos < 0 {
			break
		}

		total := a[pos] + b[pos] + carry
		if total > 9 {
			r[pos] = total - 10
			carry = 1
		} else {
			r[pos+1] = total
			carry = 0
		}
	}
	r[0] = carry
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

func DivideV1(a int, b int) (int, int) {
	q, r := 0, a
	for {
		if r < b {
			break
		}

		r = r - b
		q = q + 1
	}

	return q, r
}

func QuickDivision(a int, b int) (int, int) {

	counter, power, mid, appr := 0, 1, 0, 0

	for {
		if power*b > a {
			break
		}

		counter++
		power = power * 2
	}
	p, q := power, power/2

	for k := 1; k < counter; k++ {
		comp := (p + q) / 2
		mid = comp * b

		if mid <= a {
			appr = mid
			q = comp
		} else {
			p = comp
		}
	}

	r := a - appr
	return q, r
}

// https://blog.csdn.net/gd007he/article/details/69055031
// https://blog.csdn.net/gd007he/article/details/68961974
func PostiveSub(a []int, b []int, r []int) {
	// assuem a > b
	for i := 0; i < len(r); i++ {
		r[i] = 0
	}

	al, bl := NumbericLength(a), NumbericLength(b)

	c := 0
	j := len(r) - 1
	pos, pa, pb := 0, 0, 0

	for {
		pos = len(r) - 1 - j
		pa = len(a) - 1 - pos
		pb = len(b) - 1 - pos

		if c >= al && c >= bl {
			break
		}

		if pa >= 0 && pb >= 0 {
			if a[pa] >= b[pb] {
				r[j] = a[pa] - b[pb]
			} else {
				r[j] = a[pa] + 10 - b[pb]
				for k := pa - 1; k >= 0; k-- {
					if a[k] == 0 {
						a[k] = 9
						r[j-(pa-k)] = 9
					} else {
						tmp := a[k] - 1
						r[j-(pa-k)] = tmp
						a[k] = tmp
						break
					}
				}
			}
		} else if pb < 0 {
			r[j] = a[pa]
		}

		c++
		j--
	}
}

// http://fredrikj.net/blog/2011/03/100-mpmath-one-liners-for-pi/
// Calculate PI

// decimal implements arbitrary precision, decimal floating-point numbers, per the General Decimal Arithmetic specification.
// go get github.com/ericlagergren/decimal

func Quo(a []int, b []int, r []int) {

}

func Prime(n int) bool {
	i := 2
	for {
		if i*i > n {
			break
		}

		if n%i == 0 {
			return false
		}

		i++
	}
	return true
}

func DivisorV2(n int) int {

	num, i := 0, 1
	for {
		if i*i > n {
			break
		}

		if n%i == 0 {
			num += 2
		}
		i += 1
	}

	if i*i == n {
		num -= 1
	}

	return num
}

func DivisorV3(s int) []int {

	n, i := s, 0

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	res := make([]int, 0)

	for {
		if primes[i]*primes[i] > s {
			break
		}

		for n%primes[i] == 0 {
			res = append(res, primes[i])
			n = n / primes[i]
		}

		i++
	}

	if n < s && n > 1 {
		res = append(res, n)
	}

	return res
}

const float64EqualityThreshold = 1e-9

func FloatEqualTo(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold*(math.Abs(a)+math.Abs(b))
}

func FloatCompare(a, b float64) int {
	if FloatEqualTo(a, b) {
		return 0
	}

	if a > b {
		return 1
	} else {
		return -1
	}
}
