func productExceptSelf(nums []int) []int {
	n := len(nums)
	prex := make([]int , n)
	suff := make([]int, n)
	res := make([]int, n)
	
	prex[0] = 1
	suff[n-1] = 1
	for i := 1; i < n; i++ {
		prex[i] = nums[i-1] * prex[i-1]
	}

	for i := n-2; i >= 0; i-- {
		suff[i] = nums[i+1] * suff[i+1]
	}

	for i := 0; i < n; i++ {
		res[i] = prex[i] * suff[i]
	}
	return res
}
