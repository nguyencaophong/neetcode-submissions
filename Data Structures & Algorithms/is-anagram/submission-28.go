
func isAnagram(s string, t string) bool {
    sMap := make(map[rune]int, len(s))

    if len(s)!= len(t) {
        return false
    }

    for _, r := range s {
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
