package math

// https://leetcode.com/problems/ugly-number/
func isUgly(num int) bool {
	factors := []int{2, 3, 5}
	idx := 0
	for num > 1 {
		if idx >= len(factors) {
			break
		}
		if num%factors[idx] == 0 {
			num = num / factors[idx]
		} else {
			idx += 1
		}
	}
	return 1 == num
}
