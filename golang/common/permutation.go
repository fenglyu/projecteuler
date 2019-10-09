package common

import (
	//"testing"
	"fmt"
)

func NextPermutation(c []int, l int, r int) { //bool {
	//func NextPermutation(c interface{}, l int, r int) {

	fmt.Println("----> TODO, Support terminate")
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

	//	return false
}
