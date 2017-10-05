package hash

// https://leetcode.com/problems/distribute-candies/description/
func distributeCandies(candies []int) int {
	cnt := len(candies)
	m := make(map[int]int, 0)
	for _, k := range candies {
		m[k] += 1
	}

	kinds := 0
	for kind, ccnt := range m {
		if kinds >= cnt/2 {
			break
		}
		if ccnt == 1 {
			kinds += 1
			m[kind] = 0
		}
	}
	curCandies := kinds
	hits := make(map[int]bool)
	for curCandies < cnt/2 {
		for kind, ccnt := range m {
			if ccnt <= 0 {
				continue
			}
			if hits[kind] || curCandies >= cnt/2 {
				return kinds
			}
			hits[kind] = true
			kinds += 1
			curCandies += 1
			m[kind] -= 1
		}
	}
	return kinds
}
