package bits

// https://leetcode.com/problems/single-number/description/

func singleNumber(nums []int) int {
	cnt := len(nums)
	if cnt == 1 {
		return nums[0]
	}
	ret := nums[0]
	for i := 1; i < cnt; i++ {
		ret = ret ^ nums[i]
	}
	return ret
}
