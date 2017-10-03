package math

// https://leetcode.com/problems/set-mismatch/description/
func findErrorNums(nums []int) []int {
	ret := []int{0, 0}
	m := make(map[int]int, 0)
	for _, num := range nums {
		m[num] += 1
	}

	cnt := len(nums)
	for n := 1; n <= cnt; n++ {
		if m[n] == 1 {
			continue
		} else if m[n] == 0 {
			ret[1] = n
		} else {
			ret[0] = n
		}
	}
	return ret
}
