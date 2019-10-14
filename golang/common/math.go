package common

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

/*
func Divide(n int) int {

}
*/
