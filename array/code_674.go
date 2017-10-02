package array

// https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/
func findLengthOfLCIS(nums []int) int {
	cnt := len(nums)
	if cnt <= 1 {
		return cnt
	}
	ret := 1
	clen := 1 // 当前子串长度
	for i := 1; i < cnt; i++ {
		if nums[i] <= nums[i-1] {
			clen = 1
			continue
		}
		clen += 1
		if clen > ret {
			ret = clen
		}
	}
	return ret
}
