package array

// https://leetcode.com/problems/longest-increasing-subsequence/description/
func lengthOfLIS(nums []int) int {
	cnt := len(nums)
	if cnt <= 1 {
		return cnt
	}
	dp := make([]int, cnt+1)
	dp[1] = nums[0]
	clen := 1
	for i := 1; i < cnt; i++ {
		if nums[i] > dp[clen] {
			clen += 1
			dp[clen] = nums[i]
			continue
		}
		idx := bSearch(dp, 1, clen, nums[i])
		dp[idx] = nums[i]
	}
	return clen
}

func bSearch(dp []int, l, r, dst int) int {
	for l <= r {
		m := (l + r) / 2
		if dp[m] == dst {
			return m
		} else if dp[m] > dst {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
