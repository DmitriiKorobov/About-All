package main

import "testing"

func TestTwoSum(t *testing.T) {
	// Arrange
	testTable := []struct {
		nums []int
		target int
		expected []int
	} {
		{
			nums: []int{3, 2, 4},
			target: 6,
			expected: []int{1, 2},
		},
		{
			nums: []int{3,3}
			target: 6
			expected: []int{0, 1},
		},
		{
			nums: []int{2,7,11,15}
			target: 9
			expected: []int{0, 1},
		},
	}

	// Act
	for _, testCase := range testTable {
		result := twoSum(testCase.nums)
		t.Logf("Testing nums = %v, result = %d\n", testCase.nums, result)

	// Assert
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %d, got %d", testCase.expected, result)
		}
}
