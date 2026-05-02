func lengthOfLongestSubstring(s string) int {
	left, right := 0,0
	res := 0
	n := len(s)
	seen := make(map[byte]int,n)

	for right < n {
		if idx,ok := seen[s[right]]; ok && idx >= left{
			left = idx +1
		}
		
		seen[s[right]] = right
		if right - left + 1 > res {
			res = right - left +1
		}
		right++
	}
	return res
}