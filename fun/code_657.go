package fun

// https://leetcode.com/problems/judge-route-circle/description/
func judgeCircle(moves string) bool {
	x := 0
	y := 0
	for _, m := range moves {
		switch m {
		case 'U', 'u':
			y += 1
		case 'D', 'd':
			y -= 1
		case 'R', 'r':
			x += 1
		case 'L', 'l':
			x -= 1
		}
	}
	return x == 0 && y == 0
}
