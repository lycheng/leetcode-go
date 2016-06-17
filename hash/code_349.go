package hash

// https://leetcode.com/problems/intersection-of-two-arrays/
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	set := make(map[int]bool)
	for _, n := range nums1 {
		set[n] = true
	}
	for _, n := range nums2 {
		if !set[n] {
			continue
		}
		res = append(res, n)
		set[n] = false
	}
	return res
}
