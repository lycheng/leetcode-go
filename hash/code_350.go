package hash

// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/
func intersect(nums1 []int, nums2 []int) []int {

	var lessNums []int
	var moreNums []int

	if len(nums1) > len(nums2) {
		lessNums = nums2
		moreNums = nums1
	} else {
		lessNums = nums1
		moreNums = nums2
	}
	numCount := make(map[int]int, 0)

	for _, n := range lessNums {
		numCount[n] = numCount[n] + 1
	}

	result := make([]int, 0)
	for _, n := range moreNums {
		if numCount[n] <= 0 {
			continue
		}
		result = append(result, n)
		numCount[n] = numCount[n] - 1
	}
	return result
}
