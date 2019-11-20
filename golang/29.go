package main

/*

Consider all integer combinations of ab for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:

2**2=4,  2**3=8,   2**4=16,  2**5=32
3**2=9,  3**3=27,  3**4=81,  3**5=243
4**2=16, 4**3=64,  4**4=256, 4**5=1024
5**2=25, 5**3=125, 5**4=625, 5**5=3125
If they are then placed in numerical order, with any repeats removed, we get the following sequence of 15 distinct terms:

4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

How many distinct terms are in the sequence generated by ab for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?
*/

import (
	"fmt"
	//	"math"
)

func prime(n int) bool {
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

func main() {

	//	a := math.Pow(2, 3)
	//	fmt.Println(a)

	ar := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	//	fmt.Println(ar)

	//c := 0
	//for i := 2; i < 101; i++ {
	//	if prime(i) {
	//		//			fmt.Println(i)
	//		ar[c] = i
	//		c++
	//	}
	//}

	//	a, b := 2, 2

	mul := 1
	for i := 0; i < len(ar); i++ {

		mul = ar[i]

		for k := 0; k < len(ar); k++ {

			t := make([]int, 20)
			for q, _ := range t {
				t[q] = 0
			}
			c := 0
			t[c] = ar[i]
			c++

			j := k
			for {

				if j >= len(ar) {
					break
				}

				mul = mul * ar[j]
				t[c] = ar[j]

				//			fmt.Println(mul, ar[i], ar[j])

				if mul > 100 {
					break
				}

				//fmt.Println(ar[i])
				fmt.Println(t)
				c++
			}
		}
	}

}
