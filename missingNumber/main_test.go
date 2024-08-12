package main

import (
	"testing"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "Missing number is in the middle",
			numbers:  []int{1, 2, 4, 5, 6},
			expected: 3,
		},
		{
			name:     "Missing number is the first one",
			numbers:  []int{2, 3, 4, 5, 6},
			expected: 1,
		},
		{
			name:     "Missing number is the last one",
			numbers:  []int{1, 2, 3, 4, 5},
			expected: 6,
		},
		{
			name:     "Only one number, missing is 2",
			numbers:  []int{1},
			expected: 2,
		},
		{
			name:     "No numbers, missing is 1",
			numbers:  []int{},
			expected: 1,
		},
		{
			name:     "Large input with missing number",
			numbers:  []int{1, 2, 3, 4, 5, 6, 7, 9, 10},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := missingNumber(tt.numbers)
			if result != tt.expected {
				t.Errorf("missingNumber(%v) = %d; expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}
