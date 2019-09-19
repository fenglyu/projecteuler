package main

import (
	"fmt"
)

func mul(a interface{}, b interface{}) []int {

	c := []int{2, 3, 4, 5}

	f := a.([]int)
	g := b.([]int)
	fmt.Println("f: ", f, "g: ", g)

	result := make([]int, (len(f)+len(g))*2)

	d := 0
	for i, v := range f {
		for j, p := range g {
			//			result[i] = v + p
			result[i+j] = v + p
			d++
		}
	}

	fmt.Println(result)

	return c
}

func main() {
	a := []int{2, 3, 4, 5, 6}
	b := []int{4, 5, 3, 1}

	result := mul(a, b)
	fmt.Println(result)
}
