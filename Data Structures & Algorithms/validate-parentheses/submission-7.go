func isValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)

		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			if (c == ')' && top != '(') ||
				(c == '}' && top != '{') ||
				(c == ']' && top != '[') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}