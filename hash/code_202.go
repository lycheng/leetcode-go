package hash

import (
	"fmt"
	"strconv"
)

func isHappy(n int) bool {
	set := make(map[int]bool)
	for {
		if set[n] {
			break
		}
		set[n] = true
		n = sumOfSquareOfDigits(n)
	}
	return n == 1
}

func sumOfSquareOfDigits(n int) int {
	s := fmt.Sprintf("%d", n)

	sum := 0
	for _, d := range s {
		digit, _ := strconv.ParseInt(string(d), 10, 32)
		sum += int(digit) * int(digit)
	}
	return sum
}
