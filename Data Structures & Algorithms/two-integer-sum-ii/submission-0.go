func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	mapSum := make(map[int]int, n)
	
	for i,v := range numbers {
		mapSum[v] = i+1
	}

	
	res := []int{}
	for i,v := range numbers {
		targetTwoIndex := target - v
		value, ok := mapSum[targetTwoIndex]
		if ok {
			res = append(res,i+1, value)
			return res
		}
	}
	return res
}
