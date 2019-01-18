package bits

import "fmt"

// https://leetcode.com/problems/hamming-distance/
func hammingDistance(x int, y int) int {
	xx := fmt.Sprintf("%032b", x)
	yy := fmt.Sprintf("%032b", y)

	n := 0
	for i := 0; i < 32; i++ {
		if xx[i] == yy[i] {
			continue
		}
		n++
	}
	return n
}

func hammingDistance1(x int, y int) int {

	// slower
	z := x ^ y
	s := fmt.Sprintf("%b", z)

	n := 0
	for _, i := range s {
		if string(i) == "0" {
			continue
		}
		n++
	}
	return n
}
