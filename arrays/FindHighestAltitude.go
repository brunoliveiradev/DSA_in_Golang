package arrays

// HighestAltitude returns the highest altitude reached during a road trip.
// The road trip consists of n + 1 points at different altitudes. The biker starts his trip on point 0 with altitude equal to 0.
//
// Parameters:
//   - gain: An integer array of length n where gain[i] is the net gain in altitude between points i and i + 1 for all (0 <= i < n).
//
// Returns:
//
//	The highest altitude of a point during the trip.
//
// Time Complexity: O(n), where n is the length of the gain array.
// Space Complexity: O(1), as the space used does not depend on the input size.
// Used PrefixSum technique to solve.
func HighestAltitude(gain []int) int {
	highestAltitude, currentAltitude := 0, 0

	for _, altitudeGain := range gain {
		currentAltitude += altitudeGain
		if currentAltitude > highestAltitude {
			highestAltitude = currentAltitude
		}
	}

	return highestAltitude
}
