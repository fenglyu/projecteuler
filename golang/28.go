package main

import (
	"fmt"
	//	"github.com/fenglyu/projecteuler/golang/common"
)

/*
Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
*/

var LEN int = 101
var Demo [101][101]int

func square() int {

	return 1
}

func main() {

	for i := 0; i < len(Demo); i++ {
		for j := 0; j < len(Demo[i]); j++ {

			Demo[i][j] = 0
		}
	}
	fmt.Println(LEN)
	//	fmt.Println(Demo)

	x, y := LEN/2, LEN/2
	//	fmt.Println(init)

	i, w := 1, 1

	for {

		if i >= LEN*LEN {
			break
		}
		Demo[x][y] = i

		// 1. 1 --> 2
		y++
		i++
		Demo[x][y] = i

		// 2. up -> down
		for t := 0; t < w; t++ {
			x++
			i++
			Demo[x][y] = i
		}

		// 3. right -> left
		for t := 0; t < w*2; t++ {
			y--
			i++
			Demo[x][y] = i
		}

		// 4. down -> up
		for t := 0; t < w*2; t++ {
			x--
			i++
			Demo[x][y] = i
		}

		// 5. left -> right
		for t := 0; t < w*2; t++ {
			y++
			i++
			Demo[x][y] = i
		}

		w++
	}

	fmt.Println(Demo)
}
