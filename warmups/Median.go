package warmups

import "DSA/techniques"

// FindMedian calculates the median of an array of integers.
//
// The median is the middle value in a sorted list of numbers. If the list has an odd number of elements,
// the median is the element at the center. If the list has an even number of elements, the median is the
// average of the two central elements.
//
// Parameters:
//
//	arr []int32: An unsorted slice of integers.
//
// Returns:
//
//	int32: The median value of the input slice.
//
// Time Complexity: O(n log n), where n is the length of the input slice due to sorting.
// Space Complexity: O(n) for the sorting algorithm, as is use merge sort which requires additional space for the temporary arrays.
func FindMedian(arr []int32) int32 {
	n := len(arr)
	if n == 1 {
		return arr[0] // Return the only element if the array has one element
	}

	techniques.MergeSort(arr)

	mid := n / 2
	if n%2 == 0 {
		return (arr[mid-1] + arr[mid]) / 2
	}
	return arr[mid]
}
