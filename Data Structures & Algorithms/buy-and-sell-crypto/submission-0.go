func maxProfit(prices []int) int {
	left, right := 0,1
	n := len(prices)
	max := 0
	for right < n {
		if prices[left] < prices[right] {
			result := prices[right] - prices[left]
			if result > max {
				max = result
			}
		}else{
			left = right
		}
		right++
	}
	return max;
}

