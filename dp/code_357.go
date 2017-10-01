package dp

import "math"

// https://leetcode.com/problems/count-numbers-with-unique-digits/
func countNumbersWithUniqueDigits(n int) int {
	m := int(math.Min(float64(n), 10.0))
	seq := make([]int, m+1)
	seq[0] = 1

	for i := 1; i <= m; i++ {
		seq[i] = 9

		for j := 9; j >= 9-i+2; j-- {
			seq[i] *= j
		}
	}
	result := 0
	for _, num := range seq {
		result += num
	}
	return result
}
