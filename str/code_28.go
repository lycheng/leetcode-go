// package str
package str

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}

	lh := len(haystack)
	ln := len(needle)
	for i := 0; i < lh-ln+1; i++ {
		j := 0
		for ; j < ln; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == ln {
			return i
		}
	}
	return -1
}
