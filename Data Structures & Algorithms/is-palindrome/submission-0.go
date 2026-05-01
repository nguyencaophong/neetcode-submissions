func filterAlphanumeric(s string) string {
    return strings.Map(func(r rune) rune {
        if (r >= '0' && r <= '9') ||
            (r >= 'a' && r <= 'z') ||
            (r >= 'A' && r <= 'Z') {
            return r
        }
        return -1 // drop the character
    }, s)
}

func isPalindrome(s string) bool {
	input := strings.ToLower(filterAlphanumeric(s))
	n := len(input)
	lastIndex := n-1
	for i := 0; i < n/2; i++ {
		if input[i] != input[lastIndex]{
			return false
		}
		lastIndex -=1
	}
	return true
}