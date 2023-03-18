package datastructures

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{2, 3, 4, 10, 40, 50, 60, 70, 80, 90}, 10, 3},
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 7, 3},
		{[]int{1, 4, 6, 8, 10, 12, 14, 16, 18, 20}, 6, 2},
		{[]int{-20, -10, 0, 10, 20, 30, 40, 50, 60, 70}, 50, 7},
		{[]int{-5, 0, 5, 10, 15, 20, 25, 30, 35, 40}, 100, -1},
		{[]int{-25, -20, -15, -10, -5, 0, 5, 10, 15, 20}, -15, 2},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 3},
		{[]int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}, 700, 6},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			left, right := 0, len(tc.arr)-1
			result := binarySearch(tc.arr, tc.target, left, right)
			if result != tc.expected {
				t.Errorf("Expected index %d, but got %d", tc.expected, result)
			}
		})
	}
}
