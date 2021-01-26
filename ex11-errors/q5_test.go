package main

import (
	"testing"
)


func TestSum(t *testing.T) {
	x := sum(1, 2)
	if x != 3 {
		t.Errorf("Sum is not working, Expected 3 but got %v", x)
	}
}
