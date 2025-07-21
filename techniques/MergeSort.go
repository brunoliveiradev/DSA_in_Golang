package techniques

import "math/rand"

// GenerateRandomArray creates a slice of random int32 numbers.
func GenerateRandomArray(lenOfArray int) []int32 {
	if lenOfArray <= 0 {
		return []int32{}
	}

	numbers := make([]int32, lenOfArray)
	for i := 0; i < lenOfArray; i++ {
		numbers[i] = int32(rand.Intn(10000000))
	}
	return numbers
}

// MergeSort sorts the array in place using the Merge Sort algorithm.
func MergeSort(arr []int32) {
	mergeSort(arr)
}

// mergeSort recursively splits and sorts the array.
func mergeSort(inputArray []int32) {
	inputLength := len(inputArray)
	if inputLength < 2 {
		return // already sorted
	}

	// divide the array in two new arrays in the middle
	midIndex := inputLength / 2
	leftHalf := make([]int32, midIndex)
	rightHalf := make([]int32, inputLength-midIndex) // if the number is odd works this way

	// Copy elements to left and right halves, the copy function can replace the for loop bellow
	//copy(leftHalf, inputArray[:midIndex]) // exclusive, from 0 to midIndex - 1
	//copy(rightHalf, inputArray[midIndex:]) // inclusive, from midIndex till the end

	// copy elements from inputArray to the leftHalf array
	for i := 0; i < midIndex; i++ {
		leftHalf[i] = inputArray[i]
	}

	// copy elements from inputArray to the rightHalf array
	for i := midIndex; i < inputLength; i++ {
		rightHalf[i-midIndex] = inputArray[i]
	}

	// Recursively sort both halves
	mergeSort(leftHalf)
	mergeSort(rightHalf)

	// Merge sorted halves back into the original array
	merge(inputArray, leftHalf, rightHalf)
}

// merge combines two sorted slices into the original array in ascending order.
func merge(inputArray []int32, leftHalf []int32, rightHalf []int32) {
	leftSize := len(leftHalf)
	rightSize := len(rightHalf)

	// i for leftHalf, j for rightHalf, k for inputArray
	var i, j, k int

	// o objetivo deste for é comparar os elementos de ambos os "half" array, e ir inserindo em ordem ASC do menor para o maior
	// Merge elements from leftHalf and rightHalf in order
	for i < leftSize && j < rightSize {
		// compare the elements, and add the smallest element to the mergedArray
		if leftHalf[i] <= rightHalf[j] {
			inputArray[k] = leftHalf[i]
			i++
		} else {
			inputArray[k] = rightHalf[j]
			j++
		}
		k++
	}

	// Copy any remaining elements from leftHalf
	for i < leftSize {
		inputArray[k] = leftHalf[i]
		i++
		k++
	}

	// Copy any remaining elements from rightHalf
	for j < rightSize {
		inputArray[k] = rightHalf[j]
		j++
		k++
	}
}
