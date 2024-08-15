package mincoins

import (
	"slices"
	"testing"
)

// Old version test
func TestMincoins(test *testing.T) {
	tests := []struct {
		input []int
		sum   int

		expected []int
	}{
		{[]int{}, 15, []int{}},
		{[]int{1, 5, 10}, 13, []int{10, 1, 1, 1}},
		{[]int{1, 1, 1, 5, 5, 10}, 13, []int{10, 1, 1, 1}},
		{[]int{1, 1, 1, 5, 5, 12}, 13, []int{12, 1}},
		// {[]int{9, 1, 2, 3, 4, 5}, 17, []int{9, 5, 3}},
	}

	for i, t := range tests {
		res := MinCoins(t.sum, t.input)
		if !slices.Equal(res, t.expected) {
			test.Errorf("[%v] Test failed (original: %v, result: %v)", i, t.expected, res)
			return
		}
	}
}

// New version test
func TestMincoins2(test *testing.T) {
	tests := []struct {
		input []int
		sum   int

		expected []int
	}{
		{[]int{}, 15, []int{}},
		{[]int{1, 5, 10}, 13, []int{10, 1, 1, 1}},
		{[]int{1, 1, 1, 5, 5, 10}, 13, []int{10, 1, 1, 1}},
		{[]int{1, 1, 1, 5, 5, 12}, 13, []int{12, 1}},
		{[]int{9, 1, 2, 3, 4, 5}, 17, []int{9, 5, 3}},
		{GenerateSlice(100000, 1), 10000000, GenerateSlice(10000000, 1)},
	}

	for i, t := range tests {
		res := MinCoins2(t.sum, t.input)
		if !slices.Equal(res, t.expected) {
			test.Errorf("[%v] Test failed (original: %v, result: %v)", i, t.expected, res)
			return
		}
	}
}

// Benchmark test 367.7 nanosec per iteration
func BenchmarkCoins(b *testing.B) {
	coins, sum := []int{1, 5, 10}, 13
	for i := 0; i < b.N; i++ {
		MinCoins2(sum, coins)
	}
}
