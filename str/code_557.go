package str

// https://leetcode.com/problems/reverse-words-in-a-string-iii/description/
// 手动处理字符串非常慢
func reverseWords(s string) string {
	if s == "" {
		return s
	}
	end := len(s)
	var words []string

	l := 0
	r := 0
	for r < end {
		if s[r] != ' ' {
			r++
			continue
		}
		words = append(words, s[l:r])
		r++
		l = r
	}
	words = append(words, s[l:r])
	result := ""

	for idx, word := range words {
		for i := len(word) - 1; i >= 0; i-- {
			result += string(word[i])
		}
		if idx == len(words)-1 {
			break
		}
		result += " "
	}
	return result
}
