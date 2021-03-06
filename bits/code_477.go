package bits

// https://leetcode.com/problems/total-hamming-distance/description/
func totalHammingDistance(nums []int) int {
	count := len(nums)
	result := 0
	for i := 0; i < 32; i++ {
		oneCnt := 0
		for _, num := range nums {
			if (num & (1 << uint(i))) > 0 {
				oneCnt++
			}
		}
		result += oneCnt * (count - oneCnt)
	}
	return result
}
