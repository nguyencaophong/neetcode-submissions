func isAnaGrams(s string, t string) bool {
	if len(t) != len(t) {
		return false
	}

	sMap := make(map[rune]int, len(t))
	for _,r := range s {
		sMap[r]++
	}
	
	for _, r := range t {
        if sMap[r] == 0 {
            return false
        }
        sMap[r]--
    }
    return true
}

func groupAnagrams(strs []string) [][]string {

	result := make([][]string, 0)
	flatItemSetted := make(map[string]bool, len(strs))

	for i := 0; i < len(strs); i++ {

		if flatItemSetted[strs[i]] {
			continue
		}

		subResult := []string{strs[i]}
		flatItemSetted[strs[i]] = true

		for j := i + 1; j < len(strs); j++ {

			if len(strs[i]) == len(strs[j]) {

				isAnaGram := isAnaGrams(strs[i], strs[j])

				if isAnaGram {
					flatItemSetted[strs[j]] = true
					subResult = append(subResult, strs[j])
				}

			}
		}

		result = append(result, subResult)

	}

	return result
}
