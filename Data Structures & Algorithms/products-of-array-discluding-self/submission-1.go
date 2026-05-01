func productExceptSelf(nums []int) []int {
	// Division 
	// If array has two item zeros -> values of array is full zeros 
	// If array has one item zeros -> array full zeros except items at value where index zero

	n := len(nums)
	res := make([]int, n)
	zeroCounts := 0
	totalOfItems := 1
	for _, v := range nums {
		if v == 0 {
			zeroCounts++
		}else{
			totalOfItems *= v
		}
	}

	if zeroCounts > 1 {
		return res
	}

	if zeroCounts == 1 {
		for i, v := range nums {
			if v != 0 {
				res[i] = 0
			}else{
				res[i] = totalOfItems
			}
		}
		return res
	}

	for i,v := range nums {
		res[i] = totalOfItems / v
	}
	return res
}
