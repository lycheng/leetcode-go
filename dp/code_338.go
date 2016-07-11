package dp

/*
 * 如果一个数是 2 的 N 次幂，则位数为 1
 * 否则其等于小于它的最大的 2^N + x
 * x 范围为 1 ~ (2^N - 1)，直到 2^(N+1) 之后继续循环
 */
func countBits(num int) []int {
	res := make([]int, num+1)

	res[0] = 0

	p := 1
	pre := 1
	for i := 1; i <= num; i++ {
		if i == p {
			res[i] = 1
			p *= 2
			pre = 1
		} else {
			res[i] = res[pre] + 1
			pre += 1
		}
	}
	return res
}
