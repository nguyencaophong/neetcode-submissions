func productExceptSelf(nums []int) []int {
		res := make([]int,len(nums))
		for i := range nums {
			sub := 1	
			for k, u := range nums {
				if i == k {
					continue
				}
				sub *= u
			}
			res[i] = sub
		} 
		return res
}
