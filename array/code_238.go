package array

// TLE ????????
func productExceptSelf(nums []int) []int {

	lenNums := len(nums)
	dst := make([]int, lenNums)
	dst[0] = 1

	l := 1
	for i := 0; i < lenNums-1; i++ {
		l *= nums[i]
		dst[i+1] = l
	}

	r := 1
	for i := lenNums - 1; i >= 1; i-- {
		r *= nums[i]
		dst[i-1] *= r
	}
	return dst
}
