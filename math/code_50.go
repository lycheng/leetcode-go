// package math
// https://leetcode.com/problems/powx-n/
package math

func myPow(x float64, n int) float64 {
	var result float64 = 1
	postive := n > 0
	if !postive {
		n = -1 * n
	}
	for n != 0 {
		if n%2 == 0 {
			x = x * x
			n /= 2
		} else {
			result *= x
			n -= 1
		}
	}

	if !postive {
		result = 1.0 / result
	}
	return result
}
