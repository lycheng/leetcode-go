// package bits
// https://leetcode.com/problems/number-complement
package bits

func findComplement(num int) int {
	var f bool
	for i := 31; i >= 0; i-- {
		if (num & (1 << uint(i))) != 0 {
			f = true
		}
		if f {
			num ^= (1 << uint(i))
		}
	}
	return num
}
