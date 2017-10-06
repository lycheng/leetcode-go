package array

// https://leetcode.com/problems/max-consecutive-ones/description/
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	curCnt := 0
	for _, n := range nums {
		if n == 0 {
			curCnt = 0
			continue
		}
		curCnt += 1
		if curCnt > max {
			max = curCnt
		}
	}
	return max
}
