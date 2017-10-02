package array

import "fmt"

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

// 最长递增子串
func lis(nums []int) int {
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
	fmt.Println(dp)
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
