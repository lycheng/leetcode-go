package fun

// https://leetcode.com/problems/judge-route-circle/description/
func judgeCircle(moves string) bool {
	x := 0
	y := 0
	for _, m := range moves {
		switch m {
		case 'U', 'u':
			y++
		case 'D', 'd':
			y--
		case 'R', 'r':
			x++
		case 'L', 'l':
			x--
		}
	}
	return x == 0 && y == 0
}
