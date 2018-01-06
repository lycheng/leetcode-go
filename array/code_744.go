package array

// https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/
func nextGreatestLetter(letters []byte, target byte) byte {
	i := 0
	j := len(letters) - 1
	for i < j {
		if letters[i] > target || letters[j] <= target {
			break
		}

		mid := (i + j) / 2
		if letters[mid] <= target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return letters[i]
}
