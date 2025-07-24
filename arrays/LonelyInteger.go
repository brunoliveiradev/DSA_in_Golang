package arrays

// LonelyInteger finds the integer that appears only once in an array where every other integer appears twice.
// It uses the XOR operation to find the unique integer efficiently.
func LonelyInteger(arr []int32) int32 {
	if len(arr) == 1 {
		return arr[0]
	}

	lonely := int32(0)
	for _, v := range arr {
		lonely ^= v
	}
	return lonely
}
