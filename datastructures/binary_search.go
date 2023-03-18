package datastructures

// binarySearch is a recursive implementation of the binary search algorithm.
// arr: sorted integer slice
// target: integer value to search for
// left: integer representing the left boundary of the search interval
// right: integer representing the right boundary of the search interval
// Returns the index of the target in the array if found, otherwise -1.
func binarySearch(arr []int, target int, left int, right int) int {

	// Base case if the search interval is empty the target element is not in the array
	if left > right {
		return -1
	}

	// calculate the middle
	middle := left + (right-left)/2

	if arr[middle] == target { // target is in the middle of the array
		return middle
	} else if arr[middle] < target {
		return binarySearch(arr, target, middle+1, right) // if the middle of array is smaller than the target whe can redo the search in the index interval of middle+1 and end of array
	} else {
		return binarySearch(arr, target, left, middle-1) // if the middle of array is smaller than the target whe can redo the search in the index interval of beginning and the middle-1
	}
}
