package mincoins

import (
	"slices"
	"testing"
)

func TestMincoins(test *testing.T) {
	tests := []struct{
		input []int
		sum int

		expected []int
	}{
		{[]int{}, 15, []int{}},
		{[]int{1,5,10}, 13, []int{10, 1, 1, 1}},
		{[]int{1,1,1,5,5,10}, 13, []int{10, 1, 1, 1}},
		{[]int{1,1,1,5,5,12}, 13, []int{12, 1}},
		{[]int{9,1,2,3,4,5}, 17, []int{9, 5, 3}},
	}

	for i, t := range tests {
		res := minCoins(t.sum, t.input)
		if !slices.Equal(res, t.expected) {
			test.Errorf("[%v] Test failed (original: %v, result: %v)", i, t.expected, res)
			return
		}
	}
}

func TestMincoins2(test *testing.T) {
	tests := []struct{
		input []int
		sum int

		expected []int
	}{
		{[]int{}, 15, []int{}},
		{[]int{1,5,10}, 13, []int{10, 1, 1, 1}},
		{[]int{1,1,1,5,5,10}, 13, []int{10, 1, 1, 1}},
		{[]int{1,1,1,5,5,12}, 13, []int{12, 1}},
		{[]int{9,1,2,3,4,5}, 17, []int{9, 5, 3}},
	}

	for i, t := range tests {
		res := minCoins2(t.sum, t.input)
		if !slices.Equal(res, t.expected) {
			test.Errorf("[%v] Test failed (original: %v, result: %v)", i, t.expected, res)
			return
		}
	}
}