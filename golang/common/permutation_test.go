package common

import (
	"reflect"
	"sort"
	"testing"
)

func TestNextPermutation(t *testing.T) {

	ar := []int{5, 3, 4, 2, 1, 0, 7, 6}

	result := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 4, 3, 5},
		{1, 3, 2, 4, 5},
		{1, 3, 4, 2, 5},
		{1, 4, 2, 3, 5},
		{1, 4, 3, 2, 5},
		{2, 1, 3, 4, 5},
		{2, 1, 4, 3, 5},
		{2, 3, 1, 4, 5},
		{2, 3, 4, 1, 5},
		{2, 4, 1, 3, 5},
		{2, 4, 3, 1, 5},
		{3, 1, 2, 4, 5},
		{3, 1, 4, 2, 5},
		{3, 2, 1, 4, 5},
		{3, 2, 4, 1, 5},
		{3, 4, 1, 2, 5},
		{3, 4, 2, 1, 5},
		{4, 1, 2, 3, 5},
		{4, 1, 3, 2, 5},
		{4, 2, 1, 3, 5},
		{4, 2, 3, 1, 5},
		{4, 3, 1, 2, 5},
		{4, 3, 2, 1, 5},
	}

	sort.Ints(ar)
	NextPermutation(ar, 0, 7)

	if reflect.DeepEqual(ar, result[0]) {
		t.Errorf("Result: %d, Expected: %d", ar, result[0])
	}
	NextPermutation(ar, 0, 7)
	if reflect.DeepEqual(ar, result[1]) {
		t.Errorf("Result: %d, Expected: %d", ar, result[1])
	}

}
