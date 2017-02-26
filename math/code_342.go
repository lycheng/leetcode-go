// package math
package math

import (
	"strconv"
	"strings"
)

func isPowerOfFour(n int) bool {
	srcN := strconv.FormatInt(int64(n), 2)
	return strings.HasPrefix(srcN, "1") && strings.Count(srcN, "1") == 1 &&
		strings.Count(srcN, "0")%2 == 0
}
