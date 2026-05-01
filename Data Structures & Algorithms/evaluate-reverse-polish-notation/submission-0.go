func totalArr(arr []int, expression string) int {
    a := arr[len(arr)-2]
    b := arr[len(arr)-1]

    if expression == "+" {
        return a + b
    } else if expression == "-" {
        return a - b
    } else if expression == "*" {
        return a * b
    }
    return a / b
}

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, c := range tokens {
		intValue, err := strconv.Atoi(c);
		if err == nil  {
			stack = append(stack, intValue)
			continue
		}
		
		total := totalArr(stack, c) 
		stack = stack[:len(stack)-2]
		stack = append(stack,total)
	}
	return stack[0]
}
