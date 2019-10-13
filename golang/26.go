package main

import (
	"fmt"

	"github.com/fenglyu/projecteuler/golang/common"
)

func main() {

	nums := []int{0, 4}
	fib := common.TailFibLarge(nums)
	fmt.Println(fib)
}
