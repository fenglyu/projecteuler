package main

import (
	"fmt"
	"sort"

	"github.com/fenglyu/projecteuler/golang/common"
)

func permutatins(l interface{}) {

}

func main() {

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.Ints(nums)

	t := true
	// start from 1, because the sort array itself the first permutation
	for i := 1; i < 1000000; i++ {
		if !t {
			break
		}

		t = common.NextPermutationV3(nums, 0, len(nums))
	}

	//	fmt.Println(nums)

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum*10 + nums[i]
	}

	fmt.Println(sum)
}
