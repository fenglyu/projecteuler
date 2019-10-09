package main

import (
	"fmt"
)

func permutatins(l interface{}) {

}

func main() {

	//	lagest := 9876543210
	//
	//	sts := make([]byte, lagest+1)
	//	for i := 0; i < lagest+1; i++ {
	//		sts[i] = '0'
	//	}
	nums := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

}
