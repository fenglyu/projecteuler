package main

import (
	"testing"
)

func TestCount(t *testing.T) {
	input, expected = 11111, 5
	result := Count(input)

	if reseult != expected {
		t.Errorf("Result: %d, Expected: %d", result, expected)
	}
}
