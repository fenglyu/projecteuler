package common

import (
	"reflect"
	"sort"
	"testing"
)

var ar = []int{5, 3, 4, 2, 1}
var result = [][]int{
	{1, 2, 3, 4, 5},
	{1, 2, 3, 5, 4},
	{1, 2, 4, 3, 5},
	{1, 2, 4, 5, 3},
	{1, 2, 5, 3, 4},
	{1, 2, 5, 4, 3},
	{1, 3, 2, 4, 5},
	{1, 3, 2, 5, 4},
	{1, 3, 4, 2, 5},
	{1, 3, 4, 5, 2},
	{1, 3, 5, 2, 4},
	{1, 3, 5, 4, 2},
	{1, 4, 2, 3, 5},
	{1, 4, 2, 5, 3},
	{1, 4, 3, 2, 5},
	{1, 4, 3, 5, 2},
	{1, 4, 5, 2, 3},
	{1, 4, 5, 3, 2},
	{1, 5, 2, 3, 4},
	{1, 5, 2, 4, 3},
	{1, 5, 3, 2, 4},
	{1, 5, 3, 4, 2},
	{1, 5, 4, 2, 3},
	{1, 5, 4, 3, 2},
	{2, 1, 3, 4, 5},
	{2, 1, 3, 5, 4},
	{2, 1, 4, 3, 5},
	{2, 1, 4, 5, 3},
	{2, 1, 5, 3, 4},
	{2, 1, 5, 4, 3},
	{2, 3, 1, 4, 5},
	{2, 3, 1, 5, 4},
	{2, 3, 4, 1, 5},
	{2, 3, 4, 5, 1},
	{2, 3, 5, 1, 4},
	{2, 3, 5, 4, 1},
	{2, 4, 1, 3, 5},
	{2, 4, 1, 5, 3},
	{2, 4, 3, 1, 5},
	{2, 4, 3, 5, 1},
	{2, 4, 5, 1, 3},
	{2, 4, 5, 3, 1},
	{2, 5, 1, 3, 4},
	{2, 5, 1, 4, 3},
	{2, 5, 3, 1, 4},
	{2, 5, 3, 4, 1},
	{2, 5, 4, 1, 3},
	{2, 5, 4, 3, 1},
	{3, 1, 2, 4, 5},
	{3, 1, 2, 5, 4},
	{3, 1, 4, 2, 5},
	{3, 1, 4, 5, 2},
	{3, 1, 5, 2, 4},
	{3, 1, 5, 4, 2},
	{3, 2, 1, 4, 5},
	{3, 2, 1, 5, 4},
	{3, 2, 4, 1, 5},
	{3, 2, 4, 5, 1},
	{3, 2, 5, 1, 4},
	{3, 2, 5, 4, 1},
	{3, 4, 1, 2, 5},
	{3, 4, 1, 5, 2},
	{3, 4, 2, 1, 5},
	{3, 4, 2, 5, 1},
	{3, 4, 5, 1, 2},
	{3, 4, 5, 2, 1},
	{3, 5, 1, 2, 4},
	{3, 5, 1, 4, 2},
	{3, 5, 2, 1, 4},
	{3, 5, 2, 4, 1},
	{3, 5, 4, 1, 2},
	{3, 5, 4, 2, 1},
	{4, 1, 2, 3, 5},
	{4, 1, 2, 5, 3},
	{4, 1, 3, 2, 5},
	{4, 1, 3, 5, 2},
	{4, 1, 5, 2, 3},
	{4, 1, 5, 3, 2},
	{4, 2, 1, 3, 5},
	{4, 2, 1, 5, 3},
	{4, 2, 3, 1, 5},
	{4, 2, 3, 5, 1},
	{4, 2, 5, 1, 3},
	{4, 2, 5, 3, 1},
	{4, 3, 1, 2, 5},
	{4, 3, 1, 5, 2},
	{4, 3, 2, 1, 5},
	{4, 3, 2, 5, 1},
	{4, 3, 5, 1, 2},
	{4, 3, 5, 2, 1},
	{4, 5, 1, 2, 3},
	{4, 5, 1, 3, 2},
	{4, 5, 2, 1, 3},
	{4, 5, 2, 3, 1},
	{4, 5, 3, 1, 2},
	{4, 5, 3, 2, 1},
	{5, 1, 2, 3, 4},
	{5, 1, 2, 4, 3},
	{5, 1, 3, 2, 4},
	{5, 1, 3, 4, 2},
	{5, 1, 4, 2, 3},
	{5, 1, 4, 3, 2},
	{5, 2, 1, 3, 4},
	{5, 2, 1, 4, 3},
	{5, 2, 3, 1, 4},
	{5, 2, 3, 4, 1},
	{5, 2, 4, 1, 3},
	{5, 2, 4, 3, 1},
	{5, 3, 1, 2, 4},
	{5, 3, 1, 4, 2},
	{5, 3, 2, 1, 4},
	{5, 3, 2, 4, 1},
	{5, 3, 4, 1, 2},
	{5, 3, 4, 2, 1},
	{5, 4, 1, 2, 3},
	{5, 4, 1, 3, 2},
	{5, 4, 2, 1, 3},
	{5, 4, 2, 3, 1},
	{5, 4, 3, 1, 2},
	{5, 4, 3, 2, 1},
}

func TestNextPermutation(t *testing.T) {

	sort.Ints(ar)

	/*
		NextPermutationV1(ar, 0, len(ar)-1)
		if !reflect.DeepEqual(ar, result[0]) {
			t.Errorf("Result: %d, Expected: %d", ar, result[0])
		}

		NextPermutationV1(ar, 0, len(ar)-1)
		if !reflect.DeepEqual(ar, result[1]) {
			t.Errorf("Result: %d, Expected: %d", ar, result[1])
		}
	*/
}

func TestNextPermutationV2(t *testing.T) {
	// Sort
	sort.Ints(ar)
	flag, i := true, 0

	for {
		if i >= len(result) {
			break
		}
		//		fmt.Printf("Result: %d, Expected: %d\n", ar, result[i])
		if !reflect.DeepEqual(ar, result[i]) {
			t.Errorf("Result: %d, Expected: %d", ar, result[i])
		}
		flag = NextPermutationV2(ar, 0, len(ar)-1)
		i++
	}

	// end of iterator
	if flag {
		t.Errorf("End of iterator errof %d", ar)
	}
}

func TestNextPermutationV3(t *testing.T) {
	// Sort
	sort.Ints(ar)

	flag, i := true, 0
	for {
		if i >= len(result) {
			break
		}
		//		fmt.Printf("Result: %d, Expected: %d\n", ar, result[i])
		if !reflect.DeepEqual(ar, result[i]) {
			t.Errorf("Result: %d, Expected: %d", ar, result[i])
		}
		flag = NextPermutationV3(ar, 0, len(ar))
		i++
	}
	// end of iterator
	if flag {
		t.Errorf("End of iterator errof %d", ar)
	}
}

/*

 dlv test github.com/fenglyu/projecteuler/golang/common

 b TestNextPermutationV3

 n
*/
