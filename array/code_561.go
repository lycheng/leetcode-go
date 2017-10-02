package array

import "sort"

// https://leetcode.com/problems/array-partition-i/
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ret := 0
	for i, num := range nums {
		if i%2 == 0 {
			ret += num
		}
	}
	return ret
}
