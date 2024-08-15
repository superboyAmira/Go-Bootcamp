package ex00

import (
	"errors"
	"testing"
)

func TestPointer(t *testing.T) {
	tests := []struct {
		arr      []int
		idx      int
		expected int
		err      error
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 5, nil},
		{[]int{1, 2, 3, 4, 5}, 1, 2, nil},
		{[]int{1, 2, 3, 4, 5}, 0, 1, nil},
		{[]int{1, 2, 3, 4, 5}, 3, 4, nil},
		{[]int{1, 2, 3, 4, 5}, 10, 0, errors.New("out of range")},
		{[]int{1, 2, 3, 4, 5}, -1, 0, errors.New("negative idx")},
		{[]int{}, 1, 0, errors.New("empty slice")},
	}

	for i, test := range tests {
		res, err := GetElement(test.arr, test.idx)
		if err != nil {
			if err.Error() != test.err.Error() {
				t.Errorf("[%v] error test, expected: [%v], result: [%v]", i, test.err.Error(), err.Error())
				return
			}
		}
		if res != test.expected {
			t.Errorf("[%v] error test, expected: [%v], result: [%v]", i, test.expected, res)
			return
		}
	}
}
