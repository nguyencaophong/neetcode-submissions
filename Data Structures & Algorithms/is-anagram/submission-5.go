import "slices"

func sortedString(s string) string {
    runes := []rune(s)
    slices.Sort(runes)
    sorted := string(runes)
    return sorted
}

func isAnagram(s string, t string) bool {
    sortedS := sortedString(s)
    sortedT := sortedString(t)

    if(sortedS != sortedT) {
        return false
    }
    return true
}
