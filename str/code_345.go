package str

import (
	"strings"
)

// https://leetcode.com/problems/reverse-vowels-of-a-string/description/
func reverseVowels(s string) string {
	vowels := map[string]bool{
		"A": true, "a": true,
		"E": true, "e": true,
		"I": true, "i": true,
		"O": true, "o": true,
		"U": true, "u": true,
	}

	target := make([]string, len(s))

	i := 0
	j := len(s) - 1

	for i <= j {
		if !vowels[string(s[i])] {
			target[i] = string(s[i])
			i++
			continue
		}

		if !vowels[string(s[j])] {
			target[j] = string(s[j])
			j--
			continue
		}
		target[i], target[j] = string(s[j]), string(s[i])

		i++
		j--
	}
	return strings.Join(target, "")
}
