func mappedValueIntoMap(s string) map[string]int {
    newMap := make(map[string]int)
    for _,r := range s {
        key := string(r)

        v, exists := newMap[key]
        var increasedValue int
        if exists  {
            increasedValue = v + 1
        }else {
            increasedValue = 1
        }
        newMap[key] = increasedValue
    }
    return newMap
}

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    mappedS := mappedValueIntoMap(s)
    mappedT := mappedValueIntoMap(t)

    for k,v := range mappedS {
        valueOfKeyInMapT, exists := mappedT[k]
        if !exists || v != valueOfKeyInMapT {
            return false
        }
    }
    return true
}
