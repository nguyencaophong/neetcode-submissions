import ("maps")
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    mapS := make(map[byte]int, len(s))
    mapT := make(map[byte]int, len(t))

    for i := 0; i < len(s); i++ {
        mapS[s[i]]++
        mapT[t[i]]++
    }

    return maps.Equal(mapS, mapT)
}
