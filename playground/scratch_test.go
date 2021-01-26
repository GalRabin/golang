package main

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(1,2,3)

	if sum != 6 {
		t.Error("Got wrong result")
	}
}

func TestSumTable(t *testing.T) {
	type test struct {
		inputs []int
		answer int
	}
	tests := []test{
		{[]int{1,2,3}, 6},
		{[]int{1,2,4}, 7},
		{[]int{1,3,4}, 8},
	}

	for _, v := range tests {
		sum := Sum(v.inputs...)

		if v.answer != sum {
			t.Error("Got wrong result")
		}
	}
}
