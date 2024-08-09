package present

import (
	"day05/internal/support"
	"errors"
	"testing"
)



func TestGetNCoolestPresents(t *testing.T) {
	tests := []struct {
		name          string
		presents      []support.Present
		n             int
		expected      []support.Present
		expectedError error
	}{
		{
			name: "n greater than size of heap",
			presents: []support.Present{
				{Value: 10, Size: 1},
				{Value: 20, Size: 2},
			},
			n:             3,
			expected:      nil,
			expectedError: errors.New("n is nil or greater than size of heap"),
		},
		{
			name: "n less than or equal to zero",
			presents: []support.Present{
				{Value: 10, Size: 1},
				{Value: 20, Size: 2},
			},
			n:             0,
			expected:      nil,
			expectedError: errors.New("n is nil or greater than size of heap"),
		},
		{
			name: "n equal to size of heap",
			presents: []support.Present{
				{Value: 10, Size: 1},
				{Value: 20, Size: 2},
			},
			n: 2,
			expected: []support.Present{
				{Value: 20, Size: 2},
				{Value: 10, Size: 1},
			},
			expectedError: nil,
		},
		{
			name: "n less than size of heap",
			presents: []support.Present{
				{Value: 10, Size: 1},
				{Value: 20, Size: 2},
				{Value: 30, Size: 3},
			},
			n: 2,
			expected: []support.Present{
				{Value: 30, Size: 3},
				{Value: 20, Size: 2},
			},
			expectedError: nil,
		},
		{
			name:          "empty presents list",
			presents:      []support.Present{},
			n:             1,
			expected:      nil,
			expectedError: errors.New("n is nil or greater than size of heap"),
		},
		{
			name: "custom test case with specific input",
			presents: []support.Present{
				{Value: 5, Size: 1},
				{Value: 4, Size: 5},
				{Value: 3, Size: 1},
				{Value: 5, Size: 2},
			},
			n: 2,
			expected: []support.Present{
				{Value: 5, Size: 1},
				{Value: 5, Size: 2},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getNCoolestPresents(tt.presents, tt.n)
			if (err != nil && tt.expectedError == nil) || (err == nil && tt.expectedError != nil) || (err != nil && err.Error() != tt.expectedError.Error()) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
			if len(result) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("expected %v, got %v", tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestGrabPresents(t *testing.T) {
	tests := []struct {
		name     string
		presents []support.Present
		memory   int
		expected []support.Present
	}{
		{
			name: "Simple case with exact fit",
			presents: []support.Present{
				{Value: 5, Size: 1},
				{Value: 4, Size: 5},
				{Value: 3, Size: 1},
				{Value: 5, Size: 2},
			},
			memory: 3,
			expected: []support.Present{
				{Value: 5, Size: 1},
				{Value: 5, Size: 2},
			},
		},
		{
			name: "Case with limited memory",
			presents: []support.Present{
				{Value: 10, Size: 2},
				{Value: 5, Size: 2},
				{Value: 6, Size: 1},
			},
			memory: 2,
			expected: []support.Present{
				{Value: 10, Size: 2},
			},
		},
		{
			name: "Case where smaller items give better value",
			presents: []support.Present{
				{Value: 3, Size: 2},
				{Value: 4, Size: 3},
				{Value: 5, Size: 4},
				{Value: 3, Size: 1},
			},
			memory: 3,
			expected: []support.Present{
				{Value: 3, Size: 1},
				{Value: 3, Size: 2},
			},
		},
		{
			name: "Empty list of presents",
			presents: []support.Present{},
			memory:   5,
			expected: []support.Present{},
		},
		{
			name: "Not enough memory to fit any present",
			presents: []support.Present{
				{Value: 3, Size: 4},
				{Value: 5, Size: 5},
				{Value: 2, Size: 6},
			},
			memory: 2,
			expected: []support.Present{},
		},
		{
			name: "All presents fit exactly",
			presents: []support.Present{
				{Value: 2, Size: 1},
				{Value: 3, Size: 2},
				{Value: 4, Size: 3},
			},
			memory: 6,
			expected: []support.Present{
				{Value: 4, Size: 3},
				{Value: 3, Size: 2},
				{Value: 2, Size: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := grabPresents(tt.presents, tt.memory)

			if len(result) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("expected %v, got %v", tt.expected[i], result[i])
				}
			}
		})
	}
}