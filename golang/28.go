package main

import (
	"fmt"
)

/* Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
*/

var LEN int = 1001
var Demo [1001][1001]int

//var LEN int = 5
//var Demo [5][5]int

func square() int {

	return 1
}

func main() {

	for i := 0; i < len(Demo); i++ {
		for j := 0; j < len(Demo[i]); j++ {

			Demo[i][j] = 0
		}
	}

	x, y := LEN/2, LEN/2

	i, w := 1, 1

	Demo[x][y] = i
	// sum of the numbers on the diagonals starts
	sum := i

	for {

		if i >= LEN*LEN {
			break
		}

		// 1. 1 --> 2
		y++
		i++
		Demo[x][y] = i

		// 2. up -> down
		for t := 0; t < 2*w-1; t++ {
			x++
			i++
			Demo[x][y] = i
		}
		sum += i

		// 3. right -> left
		for t := 0; t < w*2; t++ {
			y--
			i++
			Demo[x][y] = i
		}
		sum += i

		// 4. down -> up
		for t := 0; t < w*2; t++ {
			x--
			i++
			Demo[x][y] = i
		}

		sum += i
		// 5. left -> right
		for t := 0; t < w*2; t++ {
			y++
			i++
			Demo[x][y] = i
		}

		sum += i
		w++
	}

	fmt.Println(sum)
}
