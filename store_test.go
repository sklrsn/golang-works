package main

import (
	"testing"

	"sklrsn.github.com/Robot/store"
)

func TestSum(t *testing.T) {
	sum := store.Sum(1, 2)

	if sum != 3 {
		t.Errorf("Invalid inputs")
	}
}
