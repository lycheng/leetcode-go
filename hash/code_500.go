package hash

// https://leetcode.com/problems/keyboard-row/description/
func findWords(words []string) []string {
	m1 := makeRuneMap("qwertyuiop")
	m2 := makeRuneMap("asdfghjkl")
	m3 := makeRuneMap("zxcvbnm")

	ret := make([]string, 0)
	for _, word := range words {
		if m1[rune(word[0])] && isSameLine(word, m1) {
			ret = append(ret, word)
		} else if m2[rune(word[0])] && isSameLine(word, m2) {
			ret = append(ret, word)
		} else if m3[rune(word[0])] && isSameLine(word, m3) {
			ret = append(ret, word)
		}
	}
	return ret
}

func makeRuneMap(s string) map[rune]bool {
	m := make(map[rune]bool)
	for _, r := range s {
		m[r] = true
		m[r-32] = true
	}
	return m
}

func isSameLine(word string, m map[rune]bool) bool {
	for _, r := range word {
		if !m[r] {
			return false
		}
	}
	return true
}
