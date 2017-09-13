// package str
package str

func detectCapitalUse(word string) bool {
	A := 65
	Z := 90

	first := -1
	uppercaseCnt := 0

	for idx, c := range word {
		if int(c) >= A && int(c) <= Z {
			uppercaseCnt += 1
			if first < 0 {
				first = idx
			}
		}
	}

	if uppercaseCnt == len(word) || uppercaseCnt == 0 {
		return true
	}
	if uppercaseCnt == 1 && first == 0 {
		return true
	}
	return false
}
