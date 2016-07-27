package hash

func pattern(src string) string {
	srcMap := make(map[byte]int, 0)
	array := make([]byte, 0)

	base := byte('A')
	i := 1
	for _, r := range src {
		b := byte(r)
		if srcMap[b] == 0 {
			srcMap[b] = i
			i += 1
		}
		array = append(array, base+byte(srcMap[b]))
	}
	return string(array)
}

func isIsomorphic(s string, t string) bool {
	return pattern(s) == pattern(t)
}
