func groupAnagrams(strs []string) [][]string {
	anagrams := make([][]string, 0)

	seen := make(map[[27]int][]string)

	for _, str := range strs {
		var freq [27]int
		for _, i := range str {
			freq[i-'a'] += 1
		}

		seen[freq] = append(seen[freq], str)
	}

	for _, ags := range seen {
		anagrams = append(anagrams, ags)
	}
	return anagrams
}