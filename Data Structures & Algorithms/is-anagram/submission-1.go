import "maps"
func isAnagram(s string, t string) bool {
	sChars := make(map[rune]int)
	tChars := make(map[rune]int)

	for _, sChar := range s {
		sChars[sChar] = sChars[sChar] + 1
	}

	for _, tChar := range t {
		tChars[tChar] = tChars[tChar] + 1
	}
	return maps.Equal(sChars, tChars)
}
