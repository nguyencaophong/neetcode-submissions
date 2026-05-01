func maxArea(heights []int) int {
	res := 0
	n := len(heights)

	left, right := 0, n -1
	for left < right {
		width := min(heights[left], heights[right])
		length := right - left
		
		area := width * length
		if area > res {
			res = area
		}

	 	if heights[left] <= heights[right]{
			left++
		}else {
			right--
		}
	}
	return res
}
