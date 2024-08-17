package sleepsort

import (
	"log"
	"slices"
	"testing"
)

func TestSleepSort(t *testing.T) {
	tests := []struct{
		input []int
		res chan []int
		expected []int
	}{
		{[]int{1,3,2}, make(chan []int), []int{1,2,3}},
		{[]int{ 1, 5, 3, 2, 12, 0, 4}, make(chan []int), []int{0,1,2,3,4,5,12}},
		{[]int{}, make(chan []int), []int{}},
	}

	for i, test := range tests {
		test.res = sleepsort(test.input)
		result := <-test.res
		log.Println(result)
		if !slices.Equal(result, test.expected) {
			t.Errorf("[%v] Fail orig: %v, res: %v", i, test.expected, result)
		}
	}
}