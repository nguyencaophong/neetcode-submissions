func threeSum(nums []int) [][]int {
	res := [][]int{}
    sort.Ints(nums)

	n := len(nums)
	for i := 0; i < n; i++ {
		a := nums[i]
		if a > 0 {
			break
		}

		if i > 0 && a == nums[i-1] {
			continue
		}
		
		left, right := i + 1, n - 1

		for left < right {
			subTotal := a + nums[left] + nums[right]
			if subTotal > 0 {
				right--
			} else if subTotal < 0 {
				left++
			} else{
				res = append(res, []int{a,nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left -1] {
					left++
				}
			}
		}
	}
	return res
}
