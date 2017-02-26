// package str
// https://leetcode.com/problems/reverse-string/
package str

import (
	"strings"
)

func reverseString(s string) string {
	dst := make([]string, 0)

	for i := len(s) - 1; i >= 0; i-- {
		c := string(s[i])
		dst = append(dst, c)
	}
	return strings.Join(dst, "")
}
