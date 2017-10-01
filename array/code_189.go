package array

// https://leetcode.com/problems/rotate-array/
func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	if l <= 1 || k == 0 {
		return
	}

	var tmp int
	left2rigtht := true
	if k > l/2+1 {
		k = l - k
		left2rigtht = false
	}

	for i := 1; i <= k; i++ {
		if left2rigtht {
			tmp = nums[l-1]
			for j := l - 2; j >= 0; j-- {
				nums[j+1] = nums[j]
			}
			nums[0] = tmp
		} else {
			tmp = nums[0]
			for j := 0; j < l-1; j++ {
				nums[j] = nums[j+1]
			}
			nums[l-1] = tmp
		}
	}
}
