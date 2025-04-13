package exercises

func lengthOfLongestSubstring(s string) int {
	// use sliding window to find longest substring
	charCount := make([]int, 128)
	longestSubstring, leftIndex, rightIndex := 0, 0, 0

	// until we hit end
	for rightIndex < len(s) {
		rightChar := s[rightIndex]
		charCount[rightChar] += 1

		// move left until right char isn't a duplicate
		for charCount[rightChar] > 1 {
			leftChar := s[leftIndex]
			// moving leftChar out of window; decrement the count to accomodate
			charCount[leftChar] -= 1
			leftIndex += 1
		}

		// valid substring; assign if highest that last highest
		candidateLength := rightIndex - leftIndex + 1
		if candidateLength > longestSubstring {
			longestSubstring = candidateLength
		}

		// expand window right
		rightIndex += 1
	}

	return longestSubstring
}
