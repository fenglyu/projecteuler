package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/fenglyu/projecteuler/golang/common"
)

func printExponent(powers []int) string {
	var b strings.Builder
	var c = 1
	for _, v := range powers {
		c *= v
	}
	fmt.Fprintf(&b, "%d", c)
	return b.String()
}

func equalInt(base int, powers []int) int {

	num := 0
	var sumF, baseF float64
	sumF = 1
	baseF = float64(base)

	for i, v := range powers {
		sumF = math.Pow(baseF, float64(v))
		if common.FloatCompare(sumF, 100.0) {
			break
		}

		fmt.Println(sumF, "**", printExponent(powers[i+1:]), "->", base, powers)
		baseF = sumF
		num++
	}

	return num
}

func main() {

	m := make(map[int][]int)
	for i := 2; i < 101; i++ {
		res := common.DivisorV3(i)
		if len(res) < 2 {
			continue
		}
		//		fmt.Println("num: ", i, res)
		m[i] = res
	}

	// To store the keys in slice in sorted order
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	//	fmt.Println("map: ", m)
	sum := 0
	//	for j := 2; j <= 10; j++ {
	j := 2
	for v := range keys {
		//fmt.Println(k, v)
		t := equalInt(j, m[v])
		sum += t
	}
	//}

	fmt.Println("sum: ", 99*99-sum)
}

/*

Consider all integer combinations of ab for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:
2**2=4,  2**3=8,   2**4=16,  2**5=32
3**2=9,  3**3=27,  3**4=81,  3**5=243 4**2=16, 4**3=64,  4**4=256, 4**5=1024
5**2=25, 5**3=125, 5**4=625, 5**5=3125
If they are then placed in numerical order, with any repeats removed, we get the following sequence of 15 distinct terms:

4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

How many distinct terms are in the sequence generated by ab for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?


2**(2*2), 4 ** 2
2**(2*2*2), 4 ** 4
2**(2*2*2*2), 4**8, 16**2
2**(2*2*2*2*2), 4**16, 16**8,
2**(2*2*2*2*2*2), 4**32, 16**16,

2**(3*2), 8**2
2**(3*2*2), 8**4
2**(3*2*2*2), 8**8, 64**4
2**(3*2*2*2), 8**16, 64**8
2**(3*2*2*2*2), 8**32, 64**16

2**(3*3), 8**3
2**(3*3*3), 8**9
2**(3*3*3*3), 8**27

2**(3*3*2), 8**6, 64**3
2**(3*3*2*2), 8**12, 64**6

2**(5*2), 32**2
2**(5*2*2), 32**4
2**(5*2*2*2), 32**4

2**(5*3*2)), 32**6
2**(5*3*2*2)), 32**12

2**(5*5), 32**5

2**(7*2), 4**7
2**(7*2*2), 4**16, 16**7
2**(7*2*2*2), 4**28， 16**14

2**(9*2)  4**9
2**(9*2*2) 4**18
*/
