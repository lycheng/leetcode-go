package math

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/power-of-two/description/
func isPowerOfTwo(n int) bool {
	srcN := strconv.FormatInt(int64(n), 2)
	return strings.HasPrefix(srcN, "1") && strings.Count(srcN, "1") == 1
}
