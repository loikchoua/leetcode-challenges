package longestcommonprefix

func LongestCommonPrefix(strs []string) string {
	firstWord := strs[0]
	prefix := ""
	for j := 0; j < len(firstWord); j++ {
		for index, word := range strs {
			if index > 0 {
				if j >= len(word) {
					return prefix
				} else {
					if firstWord[j] != word[j] {
						return prefix[:j]
					}
				}
			}
		}
		prefix += string(firstWord[j])
	}
	return prefix
}
