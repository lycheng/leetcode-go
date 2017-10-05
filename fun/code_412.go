package fun

import "strconv"

// https://leetcode.com/problems/fizz-buzz/description/
func fizzBuzz(n int) []string {
	var ret []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ret = append(ret, "FizzBuzz")
		} else if i%3 == 0 {
			ret = append(ret, "Fizz")
		} else if i%5 == 0 {
			ret = append(ret, "Buzz")
		} else {
			ret = append(ret, strconv.Itoa(i))
		}
	}
	return ret
}
