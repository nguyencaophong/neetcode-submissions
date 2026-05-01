func isValid(s string) bool {
    stack := []rune{}

    for _, c := range s {
        switch c {
        case '(', '[', '{':
            // push
            stack = append(stack, c)

        case ')', ']', '}':
            // stack empty → sai
            if len(stack) == 0 {
                return false
            }

            // lấy phần tử top
            top := stack[len(stack)-1]

            // check matching
            if (c == ')' && top != '(') ||
               (c == ']' && top != '[') ||
               (c == '}' && top != '{') {
                return false
            }

            // pop
            stack = stack[:len(stack)-1]
        }
    }

    // stack phải rỗng mới hợp lệ
    return len(stack) == 0
}